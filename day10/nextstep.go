package main

import "errors"

func (c cell) nextStep(prev cell) (cell, error) {
	var next cell

	switch c.char {
	case '-':
		if prev.x < c.x { // came from the left, moving to the right
			if c.x+1 >= len(maze[0]) {
				return cell{}, errors.New("x out of bound")
			}
			next.x = c.x + 1
			next.y = c.y
			next.char = maze[next.y][next.x]
		} else if prev.x > c.x { // came from the right, moving to the left
			if c.x-1 < 0 {
				return cell{}, errors.New("x out of bound")
			}
			next.x = c.x - 1
			next.y = c.y
			next.char = maze[next.y][next.x]
		} else {
			return cell{}, errors.New("same x")
		}
	case '|':
		if prev.y < c.y { // came from the top, moving down
			if c.y+1 >= len(maze) {
				return cell{}, errors.New("y out of bound")
			}
			next.y = c.y + 1
			next.x = c.x
			next.char = maze[next.y][next.x]
		} else if prev.y > c.y { // came from the bottom, moving up
			if c.y-1 < 0 {
				return cell{}, errors.New("y out of bound")
			}
			next.y = c.y - 1
			next.x = c.x
			next.char = maze[next.y][next.x]
		} else {
			return cell{}, errors.New("same y")
		}
	case 'L':
		if prev.y < c.y { // came from the top, moving right
			if c.x+1 >= len(maze[0]) {
				return cell{}, errors.New("x out of bound")
			}
			next.y = c.y
			next.x = c.x + 1
			next.char = maze[next.y][next.x]
		} else if prev.x > c.x { // came from the right, moving up
			if c.y-1 < 0 {
				return cell{}, errors.New("y out of bound")
			}
			next.y = c.y - 1
			next.x = c.x
			next.char = maze[next.y][next.x]
		} else {
			return cell{}, errors.New("wrong prev")
		}

	case 'J':
		if prev.y < c.y { // came from the top, moving left
			if c.x-1 < 0 {
				return cell{}, errors.New("x out of bound")
			}
			next.y = c.y
			next.x = c.x - 1
			next.char = maze[next.y][next.x]
		} else if prev.x < c.x { // came from the left, moving up
			if c.y-1 < 0 {
				return cell{}, errors.New("y out of bound")
			}
			next.y = c.y - 1
			next.x = c.x
			next.char = maze[next.y][next.x]
		} else {
			return cell{}, errors.New("wrong prev")
		}

	case 'F':
		if prev.y > c.y { // came from the bottom, moving right
			if c.x+1 >= len(maze[0]) {
				return cell{}, errors.New("x out of bound")
			}
			next.y = c.y
			next.x = c.x + 1
			next.char = maze[next.y][next.x]
		} else if prev.x > c.x { // came from the right, moving down
			if c.y+1 >= len(maze) {
				return cell{}, errors.New("y out of bound")
			}
			next.y = c.y + 1
			next.x = c.x
			next.char = maze[next.y][next.x]
		} else {
			return cell{}, errors.New("wrong prev")
		}

	case '7':
		if prev.y > c.y { // came from the bottom, moving left
			if c.x-1 < 0 {
				return cell{}, errors.New("x out of bound")
			}
			next.y = c.y
			next.x = c.x - 1
			next.char = maze[next.y][next.x]
		} else if prev.x < c.x { // came from the left, moving down
			if c.y+1 >= len(maze) {
				return cell{}, errors.New("y out of bound")
			}
			next.y = c.y + 1
			next.x = c.x
			next.char = maze[next.y][next.x]
		} else {
			return cell{}, errors.New("wrong prev")
		}

	default:
		return c, errors.New("wrong char")
	}
	return next, nil
}

func whatIsS(c cell) (cell, error) {
	var up, down, left, right byte
	// first, figure out what S is
	// find the neighbours
	if c.x == 0 {
		left = ' '
	} else {
		left = maze[c.y][c.x-1]
	}
	if c.x == len(maze[0])-1 {
		right = ' '
	} else {
		right = maze[c.y][c.x+1]
	}
	if c.y == 0 {
		up = ' '
	} else {
		up = maze[c.y-1][c.x]
	}
	if c.y == len(maze)-1 {
		down = ' '
	} else {
		down = maze[c.y+1][c.x]
	}

	if (up == '|' || up == 'F' || up == '7') && (down == '|' || down == 'L' || down == 'J') {
		return cell{char: '|', x: c.x, y: c.y}, nil
	}

	if (left == '-' || left == 'F' || left == 'L') && (right == '-' || right == 'J' || right == '7') {
		return cell{char: '-', x: c.x, y: c.y}, nil
	}

	if (left == '-' || left == 'F' || left == 'L') && (up == '|' || up == 'F' || up == '7') {
		return cell{char: 'J', x: c.x, y: c.y}, nil
	}

	if (right == '-' || right == 'J' || right == '7') && (up == '|' || up == 'F' || up == '7') {
		return cell{char: 'L', x: c.x, y: c.y}, nil
	}

	if (down == '|' || down == 'J' || down == 'L') && (right == '-' || right == 'J' || right == '7') {
		return cell{char: 'F', x: c.x, y: c.y}, nil
	}

	if (down == '|' || down == 'J' || down == 'L') && (left == '-' || left == 'L' || left == 'F') {
		return cell{char: '7', x: c.x, y: c.y}, nil
	}

	return cell{}, errors.New("wrong char")
}
