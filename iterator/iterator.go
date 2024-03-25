package iterator

func ConsumerArray(values []interface{}, consumer func(i int, v interface{})) {
	for idx, val := range values {
		consumer(idx, val)
	}
}

func ConsumerMap(m map[interface{}]interface{}, consumer func(k interface{}, v interface{})) {
	for key, val := range m {
		consumer(key, val)
	}
}
