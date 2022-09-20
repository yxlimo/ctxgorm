# ctxgorm

```golang
package example

import "github.com/yxlimo/ctxgorm"


type FooTable struct {
	ID uint64 `gorm:"primaryKey"`
	Name string
}

// Use Gorm Model
rows, err := ctxgorm.Model(ctx, &model.FooTable{}).ID(1, 2, 3, 4).Find()
if err != nil {
	// ...
}

// do with foo
print(rows[0].ID)


// Or just use context
var foo &Foo{}
err := ctxgorm.Ctx(ctx).ID(1).Take(foo)
if err != nil {
// ...
}
```