package ifelse

type execute func() (error, interface{})

func Invoke(e ...execute) {

}
