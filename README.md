# EventTest
EventTest is a golang library that lets you trace what your code is doing
during its entire execution. This allows you to ensure that your code is 
working as expected and that it is taking advantage of the optimizations
that you think it is.

## Example
Let's say you have two functions that both do the same thing:
- `OptimizedFunc`: faster, but can only be called
  if cond is true
- `UnoptimizedFunc`: slower, but always works

Your code currently looks like this:
```go
if cond {
	return OptimizedFunc()
}
return UnoptimzedFunc()
```

You want to ensure that `OptimizedFunc` is being called whenever it can be.
Sure, you can write unit tests to ensure the if statement is working. But
this does not necessarily give you the confidence that the whole system
is working end to end as you would expect.

With `EventTest` you can emit events to ensure thing are happening as you 
expect them to. To do this, simply add `emitEvent` calls anywhere that you
want to test. So you can modify `OptimizedFunc` and `UnoptimizedFunc` like 
so:

```go
func OptimizedFunc() {
	eventtest.EmitEvent(eventtest.NewEvent("OptimizedFunc"))
	... // Original code
}

func UnoptimizedFunc() {
	eventtest.EmitEvent(eventtest.NewEvent("UnoptimzedFunc"))
	... // Original code
}
```

Now testing that these events are easy in your e2e tests. All you have
to do is call `startListening`, and then use one of the test helpers to ensure the events
that you expect to occur have occurred. For example, this could be a template
for your e2e test.

```go
eventtest.StartListening()
defer eventtest.StopListening()

... // Standard e2e test code that runs application and ensures
... // the output matches expected.

// Now ensure we called OptimizedFunc.
eventtest.ExpectEvents([]*eventtest.Event{eventtest.NewEvent("OptimizedFunc")})
// Ensure we never called UnoptimzedFunc.
eventtest.UnexpectedEvents([]*eventtest.Event{eventtest.NewEvent("UnoptimizedFunc")})

// Clear events for next test.
eventtest.ClearEvents()
```

For a full actual example, see the `sampleprogram` directory. Note, that you cannot 
run a test that uses eventtest in parallel with any other test, because the events
the other tests are running will conflict with them. Under the hood, `StartListening`
starts up an HTTP server on port `localhost:1111` that listens for events. `EmitEvent`
and all the test helpers make API requests to this HTTP server.

This was implemented this way to be able to keep track of events through the entire system.
If you have separate processes for different services, you can still ensure that the events
from all the different processes are working as expected. 

For the API requests to work you must set `MONGOHOUSE_ENVIRONMENT` to be `local`
when testing, otherwise every call will be a noop. This ensures that in production
the `emitEvent` calls will not slow down the code.