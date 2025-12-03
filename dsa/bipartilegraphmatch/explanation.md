# Hopcroft-Karp Algorithm: Simplified

The Hopcroft-Karp algorithm finds the maximum number of matches between Workers and Jobs. It is faster than other methods because it finds multiple matches in each round.

## How it Works

The algorithm runs in **Rounds**. Each round has two steps:

### Step 1: Find Quickest Paths (BFS)
We look for the quickest way to connect an **Available Worker** to an **Available Job**.
*   We measure distance in "steps".
*   We stop searching as soon as we find the shortest possible length (e.g., 3 steps).
*   We ignore any paths that are longer than this.

### Step 2: Add Matches (DFS)
We take the paths found in Step 1 and use them to increase our total matches.
*   We call these **Improvement Paths**.
*   We **Swap** the edges along these paths to gain +1 match for each path.

---

## Key Concepts

### 1. Improvement Path (Augmenting Path)
This is a path that starts with an **Available Worker** and ends with an **Available Job**.
It alternates between **Available** and **Taken** edges.

Example: `Available Worker -> Job A == Worker B -> Available Job`

**Symbols**:
*   `->` : **Available Edge** (Not currently used in a match)
*   `==` : **Taken Edge** (Currently used in a match)

### 2. BFS Distance Logic (`dist[worker] = dist[u] + 1`)
When we are at Worker `u` and find a Job `v` that is **Taken** by another Worker `w`:
*   We follow the match: `u -> v == w`.
*   This counts as moving to the **Next Layer**.
*   So, `dist[w] = dist[u] + 1`.
    *   If `u` is at Layer 0, `w` is at Layer 1.
    *   This builds the layers: `Layer 0 -> Layer 1 -> Layer 2...`

**Meaning of `dist[u]`**:
*   It tells us **"How many steps away is Worker `u` from a Free Worker?"**
*   `dist[u] = 0`: Worker `u` IS a Free Worker.
*   `dist[u] = 1`: Worker `u` is reached from a Free Worker.
    *   **Path**: `FreeWorker -> Job (Taken by u) == Worker u`
    *   (The Free Worker connects to a Job, which is currently matched to Worker `u`)
*   `dist[u] = 2`: Worker `u` is reached from a Free Worker via another matched worker.
    *   **Path**: `FreeWorker -> Job == Worker(dist=1) -> Job == Worker u`
*   `dist[u] = Infinite`: Worker `u` cannot be reached (or is already taken/processed).

### 3. The "Swap" (Why it works)
Every Improvement Path has an odd number of edges (1, 3, 5...).
It always has **one more** available edge than taken edges.

**Example**:
Path: `B -> a == A -> b`
*   `B -> a`: Available (New)
*   `a == A`: Taken (Old)
*   `A -> b`: Available (New)
*   **Score**: 1 Taken match (`A-a`).

**After Swapping**:
We flip every edge status.
*   `B == a`: Becomes Taken (New Match!)
*   `a -> A`: Becomes Available (Old match broken)
*   `A == b`: Becomes Taken (New Match!)
*   **Score**: 2 Taken matches (`B-a`, `A-b`).

**Result**: We gained +1 match.

### 3. Why Multiple Rounds?
In each round, we only look for the **Quickest Paths** (shortest length).
*   **Round 1**: We might find many easy matches (length 1).
*   **Round 2**: Now that those easy matches are taken, we might need to look deeper (length 3) to find more matches.
*   **Round 3**: We might need even longer paths (length 5).

We repeat this until no more paths exist. This strategy is much faster than looking for any path randomly.

---

## Example Walkthrough

**Graph**:
*   Workers: **A**, **B**
*   Jobs: **a**, **b**
*   Edges: `A-a`, `A-b`, `B-a`

**Start**: No matches.

### Round 1
1.  **Find Paths**:
    *   We see `A` connects to `a` and `b`.
    *   We see `B` connects to `a`.
    *   Shortest path length is **1 step** (direct connection).
2.  **Add Matches**:
    *   We pick `A -> a`. **Match A with a**.
    *   Can we pick `B -> a`? No, `a` is now taken by `A`.
    *   End of Round 1. Matches: `(A, a)`.

### Round 2
1.  **Find Paths**:
    *   `B` is the only Available Worker.
    *   `B` connects to `a` (Taken by `A`).
    *   `A` connects to `b` (Available!).
    *   Path found: `B -> a -> A -> b`. Length is **3 steps**.
2.  **Add Matches**:
    *   We **Swap** the edges in this path.
    *   Old Match: `(A, a)` is removed.
    *   New Matches: `(B, a)` and `(A, b)` are added.
    *   End of Round 2. Matches: `(B, a), (A, b)`.

### Round 3
*   No more paths can be found.
*   **Final Result**: 2 Matches.
