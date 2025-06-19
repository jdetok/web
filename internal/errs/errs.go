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

package errs

import (
	"errors"
	"net/http"
)

type ErrInfo struct {
	Prefix string
	Msg string
}

func (e ErrInfo) Error(err error) error {
	return errors.New("*" + e.Prefix + " error: " + e.Msg + "\n**" + err.Error())
}

func HTTPErr(w http.ResponseWriter, e error) {
	http.Error(w, e.Error(), http.StatusInternalServerError)
}

func HTTPErrOld(w http.ResponseWriter, r *http.Request, e error) {
	http.Error(w, e.Error(), http.StatusInternalServerError)
}
