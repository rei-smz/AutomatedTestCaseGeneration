package Examples

import (
	"fmt"
	"math"
)

// QuadraticSolver encapsulates core methods for solving quadratic equations.
type QuadraticSolver struct {
	// ErrorThreshold is the acceptable error for Newton's Method.
	ErrorThreshold float64
}

// NewQuadraticSolver creates a new QuadraticSolver with default error threshold.
func NewQuadraticSolver() *QuadraticSolver {
	return &QuadraticSolver{
		ErrorThreshold: 1e-8,
	}
}

// SolveQuadratic solves the quadratic equation: ax^2 + bx + c = 0.
// It returns a formatted string with the roots or an error if precision is lost.
func (qs *QuadraticSolver) SolveQuadratic(a, b, c float64) (string, error) {
	discriminant := b*b - 4*a*c

	// Check for overflow or significant precision loss.
	if math.IsNaN(discriminant) || discriminant == b*b {
		return "", fmt.Errorf("not enough precision")
	}

	// For complex roots.
	if discriminant < 0 {
		sqrtVal := qs.sqrtByNewton(-discriminant)
		realPart := (-b) / (2 * a)
		imaginaryPart := sqrtVal / (2 * a)
		formattedReal := qs.formatDouble(realPart)
		formattedImaginary := qs.formatDouble(imaginaryPart)

		// Build output similar to the Java version.
		result := "x1 = "
		if formattedReal != "0" {
			result += formattedReal + " + "
		}
		if formattedImaginary != "1" {
			result += formattedImaginary
		}
		result += "i\nx2 = "
		if formattedReal != "0" {
			result += formattedReal + " - "
		} else {
			result += "-"
		}
		if formattedImaginary != "1" {
			result += formattedImaginary
		}
		result += "i"
		return result, nil
	}

	// For real roots.
	sqrtVal := qs.sqrtByNewton(discriminant)
	q := -0.5 * (b + float64(qs.sign(b))*sqrtVal)
	x1 := qs.formatDouble(q / a)
	x2 := qs.formatDouble(c / q)
	result := "x1 = " + x1
	if x1 != x2 {
		result += "\nx2 = " + x2
	}
	return result, nil
}

// sign returns 1 if b > 0, otherwise -1 (also when b equals 0).
func (qs *QuadraticSolver) sign(b float64) int {
	if b > 0 {
		return 1
	}
	return -1
}

// sqrtByNewton computes the square root of value using Newton's Method.
func (qs *QuadraticSolver) sqrtByNewton(value float64) float64 {
	if value == 0 {
		return 0
	}
	previous := (1 + value) / 2
	var result float64
	for {
		result = (previous + value/previous) / 2
		if math.Abs(previous-result) < qs.ErrorThreshold {
			break
		}
		previous = result
	}
	return result
}

// formatDouble formats a float64 such that if the value is an integer,
// no decimal part is shown.
func (qs *QuadraticSolver) formatDouble(value float64) string {
	if value == math.Floor(value) {
		return fmt.Sprintf("%.0f", value)
	}
	return fmt.Sprintf("%g", value)
}
