# lox
For building and utilizing lock banks

## Usage
```
bank := lox.NewBank()
bank.Lock("resourceA")
// nothing else can use resourceA
...
bank.Lock("resourceB")
// nothing else can use resourceA or resourceB
...
bank.Unlock("resourceB")
bank.Unlock("resourceA")
```
See the `examples` dir for more detailed examples.
