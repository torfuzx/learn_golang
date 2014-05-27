package safemap

type UpdateFunc func(interface{}, bool) interface{}

// this speicifies the action and data that and can operate
// represents a directive that operates on a map
type commandData struct {
	action 		commandAction
	key 		string
	value 		interface{}						// a interface
	result 		chan<- interface{}				// send-only
	data 		chan<- map[string]interface{}	// send-only
	updater 	UpdateFunc
}

type safeMap chan commandData	// a bidirectional channel

// returns an interface SafeMap, this function will create a channel of commandData
func New() SafeMap {
	sm := make(safeMap)	// the type of safeMap is 'chan commandData'
	go sm.run()
	return sm
}

// SafeMap is an interface, and the safeMap is an channel for transmitting commandData
type SafeMap interface {
	Insert(string, interface{})
	Delete(string)
	Find(string) (interface{}, bool)
	Len() int
	Update(string, UpdateFunc)
	Close() map[string]interface{}
}

type commandAction int

type findResult struct {
	value interface{}
	found bool
}

const (
	remove commandAction = iota		// 0 successive untyped integer
	end								// 1
	find							// 2
	insert							// 3
	length							// 4
	update							// 5
)

func (sm safeMap) Insert (key string, value interface{}) {
	sm <- commandData{
		action	: insert,
		key		: key,
		value	: value,
	}
}

func (sm safeMap) Delete (key string) {
	sm <- commandData{
		action	: remove,
		key		: key,
	}
}

func (sm safeMap) Find (key string) (value interface{}, found bool) {
	reply := make(chan interface{})
	sm <- commandData{
		action	: find,
		key		: key,
		result	: reply,
	}
	result := (<-reply).(findResult)
	return result.value, result.found
}

func (sm safeMap) Len() int {
	reply := make(chan interface{})
	sm <- commandData{
		action	: length,
		result	: reply,
	}
	return (<-reply).(int)
}

func (sm safeMap) Close() map[string] interface {} {
	reply := make(chan map[string]interface{})
	sm <-commandData{
		action : end,
		data : reply,
	}
	return <-reply
}

func (sm safeMap) Update(key string, updater UpdateFunc) {
	sm <- commandData{
		action: update,
		key : key,
		updater : updater,
	}
}

func (sm safeMap) run() {
	// the underhood storage
	store := make(map[string] interface{})
	for command := range sm {
		switch command.action {
		case insert:
			store[command.key] = command.value
		case remove:
			delete(store, command.key)
		case find:
			value, found := store[command.key]
			command.result <- findResult{value, found}
		case length:
			command.result <- len(store)
		case update:
			value, found := store[command.key]
			store[command.key] = command.updater(value, found)
		case end:
			close(sm)
			command.result <- store
		}
	}
}


