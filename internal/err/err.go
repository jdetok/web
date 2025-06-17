// EXAMPLE IMPLEMENTATION:
/*
// beginning of function ...
e := err.ErrInfo{Prefix: "database conenction error",}
...
if err := db.Ping(); err != nil {
		e.Msg = "db.Ping() failed"
		return nil, e.ConstructError(err)
	}
*/
// EXAMPLE OUTPUT:
/*
*database conenction error: db.Ping() failed
**Error 1045 (28000): Access denied for user 'x'@'0.0.0.0' (using password: YES)
 */

package err

import "errors"

type ErrInfo struct {
	Prefix string
	Msg string
}


func (e ErrInfo) Error(err error) error {
	return errors.New("*" + e.Prefix + " error: " + e.Msg + "\n**" + err.Error())
}
