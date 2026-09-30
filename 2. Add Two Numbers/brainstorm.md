first
second
temp (default = 0)

result
current <- result

while first <> nil || second <> nil:

> sum = first.val + second.val + temp
> result.val = sum % 10
> temp = sum / 10

> first = first.next
> second = second.next

> newNode
> current.Next = newNode
> current = newNode

return result
