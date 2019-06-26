CONSTANT belong to compile time

1. All basic literals are Unnamed (~Untyped) constant values - 1, 3.14, true, "hello", false
2. Named constant will replace their value in run time
3. Untyped constant may or maynot have a type
4. Constant always have to initialized to a value
    const meter int //x error
    const meter int = 100 //no error

RULES:
1. You can't initialize a constant to runtime value
    const a int = math.Pow10(2) // error as function call belongs to run time

    n := 2
    const value = n //error as variable belongs to runtime

2. We can initialize a constant using builtin function len
    const length int = len("hello") // argument to len() is a constant

Constant can be any type
                                    // can use expression
    const min int  = 1              // =1+2
    const pi float = 3.14           // =min*2.3
    const safe bool = true          // =!true
    const msg string = "accept"     // ="2.0.1"+"beta"

    no need to declare types, as types are optional.

Multiple Declaration:
    const min, max int = 1,100
    const (
        min int = 1     
        max int = 100
    )

    NOTE: constant can get their type and expression from previous type

    const (
        min int = 1     // if min int = 1+1
        max             // max int = 1 , then max will be max int = 1+1
    )

Typeless / Untyped constant - it's type is kind

    const (
		min = 1 + 1     // int type
		max = min * 3.2 // float64 type
	)


