Generate the column name code to use without hard code column name

## How to generate code
```
import "github.com/skyrocket-qy/gormx/tablename"
columnname.GenTableNamesCode(db, path)
```

## How to use
![alt text](image.png)
```
var db *gorm.DB
db.Model(table.Roles)
```
