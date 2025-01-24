package main

import "sync"

func main() {
	var mux sync.Mutex

	mux.Lock()
	// Что-то тут делаем, но доступ будет открыт только для то горутины, которая вызвала Lock.
	// Когда она сделает Unlock, доступ получит следующая горутина, первая вызвавшая Lock.
	// Данный вид синхронизации используется, когда несколько горутин читают/пишут из одной памяти.
	mux.Unlock()

	mux.TryLock() // возвращает bool, можно ли сейчас заблокировать или нет. Сама блокировка не будет произведена! Редко применяется

	// Главный недочёт обычного мьютекса в том, что он блокирует память для всех других горутин.
	// Но ведь читать из памяти можно параллельно. А вот писать должно только одна горутина и все чтения нужно остановить.
	// Для этого придумали RWMutex

	var rwmux sync.RWMutex

	rwmux.RLock() // Блокировка для горутины, которая будет читать

	rwmux.RUnlock()

	rwmux.Lock() // Блокировка для горутины, которая будет писать

	rwmux.Unlock()

	// RWMutex разделяет читающие и пишущие горутины. Если случился RLock(), то чтение из памяти не блокируется, но блокируется запись.
	// Если случился Lock(), то блокируется как чтение, так и запись.

	// Важно учитывать, что RLock() значительно быстрее обычного Lock() у Mutex. Однако Lock() у RWMutex медленней, чем у Mutex.
	// Поэтому RWMutex нужно использовать, если горутин на чтение > 70% общего числа горутин!

}
