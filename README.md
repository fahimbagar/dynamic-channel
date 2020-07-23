# Select Subset of Dynamic Size Channel

This is from stackoverflow [question](https://stackoverflow.com/questions/63003656/how-to-read-from-subset-of-channels)

- There are N channels.
- Every receiver wants to get the next message from one of the subsets of the channels.
- Need not to loose the messages and keep the messages order by the time.

Here is an example.
1. We have two channels C1 and C2.
2. There is a new message in the C2 channel. No receivers. Waiting.
3. The first receiver wants to read messages from C1 only. No messages in C1. Waiting.
4. A new message appears in the C1. The first receiver gets this message.
5. There is a new message in the C1 channel. No receivers. We have messages in C1 and C2 channels.
6. The second receiver wants to read messages from C1 and C2. The following is very important. The second receiver should get the message exactly from C2 (not a random from C1 or C2) because the message in the C2 channel appeared earlier!

## Use Reflect
Output:
```shell script
$ go run main.go                                                                                                                 12:05:50
  want from Chan CH1
  want from Chan CH1
  [CH1] waiting for 2 seconds before checking again
  want from Chan CH1
  result from CH1: 1
  want from Chan CH2
  result from CH1 & CH2: 2
  [CH1] waiting for 2 seconds before checking again
  [CH1] waiting for 2 seconds before checking again
  waiting for 5 minute before sending int to CH2
  Press CTRL+C to exit
  result from CH1: 1
  ^CExited
```