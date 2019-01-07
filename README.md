# Learning about Golang context

Prompted by:

* https://stackoverflow.com/questions/53983638/how-can-i-set-up-the-logging-context-from-middleware
* https://stackoverflow.com/questions/54016022/why-does-sharing-state-via-context-only-work-with-my-request-middleware

# Notes

## 4 won't work

Since copies of the struct are used in non-pointer receiver methods

## 5 has a race condition

https://stackoverflow.com/questions/54058007/

Run `make` to prove it has a race condition

## 7 has my favoured solution

WHERE CONTEXT WAS ABSOLUTELY UNNECESSARY!

Via @RenThraysk, which can about via a comment on my Youtube video! https://www.youtube.com/watch?v=yvQecP-8uJ8
