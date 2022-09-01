// 622. 设计循环队列
// URL：https://leetcode.cn/problems/design-circular-queue/
// 难度：中等
// 关键词：队列、数组
// 执行用时：104 ms, 在所有 TypeScript 提交中击败了 80.84% 的用户
// 内存消耗：49.5 MB, 在所有 TypeScript 提交中击败了 40.91% 的用户

class MyCircularQueue {
  #queue: Int32Array
  #cnt = 0
  #head = 0
  #rear = 0

  constructor(k: number) {
    this.#queue = new Int32Array(k)
  }

  enQueue(value: number): boolean {
    if (this.isFull())
      return false
    this.#queue[this.#rear] = value
    this.#rear = (this.#rear + 1) % this.#queue.length
    this.#cnt++
    return true
  }

  deQueue(): boolean {
    if (this.isEmpty())
      return false
    this.#head = (this.#head + 1) % this.#queue.length
    this.#cnt--
    return true
  }

  Front(): number {
    return this.isEmpty() ? -1 : this.#queue[this.#head]
  }

  Rear(): number {
    return this.isEmpty() ? -1 : this.#queue[(this.#rear + this.#queue.length - 1) % this.#queue.length]
  }

  isEmpty(): boolean {
    return this.#cnt === 0
  }

  isFull(): boolean {
    return this.#cnt === this.#queue.length
  }
}
