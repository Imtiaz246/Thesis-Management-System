package utils

func SetIfNonDefault[T comparable](src, dst *T) {
	if src == nil {
		return
	}
	var def T
	if *src != def {
		*dst = *src
	}
}
