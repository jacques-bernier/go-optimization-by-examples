Where this can be useful?

In go, hashing is done against an interface `hash.Hash`. If you need extremely high performance, you can re-implement the hashing. 
You can still achieve this without re-implementing crypto algorithms. See how segment handled this https://github.com/segmentio/fasthash 