package tasks

//----------------------------------------------
// Task payload.
//---------------------------------------------

/*
type <name>Payload struct {
	...
}
*/

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//---------------------------------------------

/*
func New<name>Task(...) (*asynq.Task, error) {
	payload, err := json.Marshal(<name>Payload{...})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(Type<name>, payload), nil
}
*/

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface.
//---------------------------------------------------------------

/*
func Handle<name>Task(ctx context.Context, t *asynq.Task) error {
	var p <name>Payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	dbClient, asynqClient := getClient(ctx)

	return nil
}
*/
