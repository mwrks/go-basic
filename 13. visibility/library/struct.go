package library

type student struct {	// unexported struct
	Name  string	// exported field
	grade int	// unexported field
}
type Student struct {	// exported struct
	Name  string	// exported field
	grade int	// unexported field
}
type StudentExp struct {	// exported struct
	Name  string	// exported field
	Grade int	// unexported field
}
