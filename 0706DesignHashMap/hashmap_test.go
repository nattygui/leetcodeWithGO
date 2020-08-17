package haspmap

import "testing"

func TestOK(t *testing.T) {
	hashmap := Constructor()

	hashmap.Put(1, 1)
	hashmap.Put(2, 2)

	value := hashmap.Get(1)
	if value != 1 {
		t.Errorf("错误, 应该为1 实际为%v", value)
	}
	value = hashmap.Get(2)
	if value != 2 {
		t.Errorf("错误, 应该为2 实际为%v", value)
	}

	hashmap.Put(2, 1)
	value = hashmap.Get(2)
	if value != 1 {
		t.Errorf("错误, 应该为1 实际为%v", value)
	}

	hashmap.Remove(2)
	value = hashmap.Get(2)
	if value != -1 {
		t.Errorf("错误, 应该为-1 实际为%v", value)
	}
}
