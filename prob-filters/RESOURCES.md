# Further Reading and Resources

This document provides a curated list of resources for deeper understanding and practical application of probabilistic data structures.

## General Resources

*   **Probabilistic Data Structures for Big Data Applications**: A comprehensive overview of various probabilistic data structures and their applications.

*   **Awesome Probabilistic Data Structures**: A GitHub repository with a curated list of probabilistic data structures, implementations, and resources.

## Model-Specific Resources

### Bloom Filter

*   **Original Paper**: Bloom, B. H. (1970). Space/time trade-offs in hash coding with allowable errors. Communications of the ACM, 13(7), 422-426.
*   **Wikipedia**: [Bloom filter](https://en.wikipedia.org/wiki/Bloom_filter)
*   **Blog Post**: [Bloom Filters by Example](https://llimllib.github.io/bloomfilter-tutorial/)

### Counting Bloom Filter

*   **Original Paper**: Fan, L., Cao, P., Almeida, W., & Broder, A. (2000). Summary Cache: A Scalable Wide-Area Web Cache Sharing Protocol. IEEE/ACM Transactions on Networking, 8(3), 281-293.
*   **Wikipedia**: [Counting Bloom filter](https://en.wikipedia.org/wiki/Counting_Bloom_filter)

### Cuckoo Filter

*   **Original Paper**: Fan, B., Andersen, D. G., Kaminsky, M., & Mitzenmacher, M. D. (2014). Cuckoo Filter: Better Than Bloom. ACM Transactions on Computer Systems (TOCS), 32(3), 1-24.
*   **Wikipedia**: [Cuckoo filter](https://en.wikipedia.org/wiki/Cuckoo_filter)

### Quotient Filter

*   **Original Paper**: Bender, M. A., Farach-Colton, M., Kuszmaul, B. C., & Zadeh, R. B. (2012). The Quotient Filter: A New Filter for Big Data. arXiv preprint arXiv:1205.0650.
*   **Blog Post**: [The Quotient Filter](https://www.cs.cmu.edu/~dga/papers/qf_blog.html)

### XOR Filter

*   **Original Paper**: D. Lemire, O. Kaser, and K. A. S. (2019). XOR Filters: Faster and Smaller Than Bloom Filters. Journal of Experimental Algorithmics, 24, 1.
*   **Blog Post**: [XOR Filters: Faster and Smaller Than Bloom Filters](https://lemire.me/blog/2019/03/19/xor-filters-faster-and-smaller-than-bloom-filters/)

### HyperLogLog

*   **Original Paper**: Flajolet, P., Fusy, É., Gandouet, O., & Meunier, F. (2007). HyperLogLog: the analysis of a near-optimal cardinality estimator. AofA.
*   **Wikipedia**: [HyperLogLog](https://en.wikipedia.org/wiki/HyperLogLog)
*   **Blog Post**: [Damn Cool Algorithms: Cardinality Estimation](https://highlyscalable.wordpress.com/2012/05/01/damn-cool-algorithms-cardinality-estimation/)

### Count-Min Sketch

*   **Original Paper**: Cormode, G., & Muthukrishnan, S. (2005). An improved data stream summary: the count-min sketch and its applications. Journal of Algorithms, 55(1), 58-74.
*   **Wikipedia**: [Count-Min Sketch](https://en.wikipedia.org/wiki/Count-Min_sketch)
*   **Blog Post**: [Count-Min Sketch: A Probabilistic Data Structure for Counting Frequencies](https://www.geeksforgeeks.org/count-min-sketch-a-probabilistic-data-structure-for-counting-frequencies/)

### t-digest

*   **Original Paper**: Dunning, T., & Ertl, O. (2014). t-digest: A new data structure for accurate quantiles on the fly. arXiv preprint arXiv:1407.0209.
*   **GitHub**: [tdigest (Java implementation)](https://github.com/tdunning/t-digest)

### Top-K

*   **General Concept**: Often implemented using algorithms like Misra-Gries or Frequent.
*   **Paper (Misra-Gries)**: Misra, J., & Gries, D. (1982). Finding repeated elements. Science of Computer Programming, 2(2), 143-152.
*   **Blog Post**: [Finding Frequent Items in Data Streams](https://www.cs.princeton.edu/courses/archive/fall12/cos521/lectures/07-frequent-items.pdf)

### Skip List

*   **Original Paper**: Pugh, W. (1990). Skip lists: a probabilistic alternative to balanced trees. Communications of the ACM, 33(6), 668-676.
*   **Wikipedia**: [Skip list](https://en.wikipedia.org/wiki/Skip_list)
*   **Blog Post**: [Skip Lists: Done Right](https://attractivechaos.wordpress.com/2013/05/20/skip-lists-done-right/)

### MinHash / Locality Sensitive Hashing (LSH)

*   **Original Paper (MinHash)**: Broder, A. Z. (1997). On the resemblance and containment of documents. In Compression and Complexity of Sequences 1997 (pp. 21-29). IEEE.
*   **Original Paper (LSH)**: Gionis, A., Indyk, P., & Motwani, R. (1999). Similarity search in high dimensions via hashing. VLDB.
*   **Blog Post**: [Understanding MinHash and Locality Sensitive Hashing](https://towardsdatascience.com/understanding-minhash-and-locality-sensitive-hashing-lsh-f62f62f62f62)

### HyperLogLog++

*   **Original Paper**: He, Y., & Manku, G. (2012). HyperLogLog in Practice: Revisiting Cardinality Estimation of Large Data Streams. VLDB.
*   **Blog Post**: [HyperLogLog++: The Algorithm for Counting Unique Elements](https://engineering.fb.com/2018/12/13/data-infrastructure/hyperloglog/)