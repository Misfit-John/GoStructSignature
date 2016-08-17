# GoStructSignature
Get a signature from the struct define. Don't care about the value, just struct.

## How to Use:
import the module

`import "github.com/johnzeng/GoStructSignature"`

And call it like:

```
signature := StructSignature.GetSignature(AnyStruct{})
```

You can put anything into GetSignature method, it will generate you a md5 string as the signature of this struct


