1. The ultimate goal of a type system is preventing bugs as early as possible

Dynamic Type:
    Postpone type checking at runtime
Static Type:
    Type checking at compile time
GO is static type language

2. Predeclared type is built-in type can use everywhere
    int,float,bool,string

    int details:
        Representation : -1,0,100,568798989
        Size :            8 bytes
    Bool details:
        Representation : true or false
        Size           : 1 byte

3. Defined Type (AKA named type)
    It can be declared from existing type (int,float,bool,string)

    eg : type Duration int64
    It's new type Duration based on underlying type int64
    int64 has representation and size -> so Duration acquire all features of underlying type

    but 
    var ms Duration = 100
    var ns int64
    ns = ms (x error due to type mismatch)

    ns = int64(ms) (Correct ...)

    Note: Use of defined type - as METHODs can only be attached with defined types


4. Aliased Type
    - byte and uint8 are same type but different names
        type byte = uint8
    - rune and int32 are same type with different names
        type rune = int32


