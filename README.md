
# Install

    go get toolman.org/base/osutil

# osutil
`import "toolman.org/base/osutil"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package osutil provides convenience functions for OS interaction




## <a name="pkg-index">Index</a>
* [func FindEnvDefault(defaultValue string, vars ...string) string](#FindEnvDefault)


#### <a name="pkg-files">Package files</a>
[env.go](/src/toolman.org/base/osutil/env.go) 





## <a name="FindEnvDefault">func</a> [FindEnvDefault](/src/target/env.go?s=382:445#L2)
``` go
func FindEnvDefault(defaultValue string, vars ...string) string
```
FindEnvDefault examines the environment variables specified by vars and
returns the value of the first variable found to be set, even if the set
value is the empty string. If none of the given vars are set in the current
environment, defaultValue is returned instead.


