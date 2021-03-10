# timedecay

sorting algorithms implemented for go, based heavily on the nodejs implementation of [https://github.com/clux/decay](https://github.com/clux/decay)

This library houses 3 popularity estimating algorithms employed by bigger news sites used to sort for best content:

  1. `WilsonScore` - Reddit's _best_ comment scoring system
  2. `RedditHot` - Reddit's _hot_ post scoring system for news posts
  3. `HackerHot` - Hackernews' scoring system

![Wilson score equation](https://github.com/clux/decay/raw/master/rating-equation.png)

## Usage
Decay exports 3 scoring function factories.

Two of these algorithms decay with time, and the other is based purely on statistical popularity.

### timedecay.NewWilsonScore(float64)
Non-decay based calculation based on statistical confidence of data.
```
import (
    "github.com/bububa/timedecay"
)

func main() {
    wilsonScore := timedecay.NewWilsonScore(1.96)
    fmt.println(wilsonScore.Score(10, 3))
}
```

### timedecay.NewRedditHot(float64)
Popularity calcuation based on published Reddit rankings.  [http://amix.dk/blog/post/19588](http://amix.dk/blog/post/19588)
```
import (
    "time"

    "github.com/bububa/timedecay"
)

func main() {
    redditHot := timedecay.NewRedditHot(45000)
    fmt.println(redditHot.Score(10, 3, time.now()))
}
```

### timedecay.NewHackerHot(float64)
Popularity calcuation based on published HackerNews rankings.  [http://amix.dk/blog/post/19574](http://amix.dk/blog/post/19574)
```
import (
    "time"

    "github.com/bububa/timedecay"
)

func main() {
    hackerHot := timedecay.NewHackerHot(1.9)
    fmt.println(hackerHot.Score(10, 3, time.now()))
}
```

## Parameter Explanation
### 1. Wilson Score
AKA Reddit's *[Best](http://blog.reddit.com/2009/10/reddits-new-comment-sorting-system.html)* comment sorting system. [Source](https://github.com/reddit/reddit/blob/bd922104b971a5c6794b199f364a06fdf61359a2/r2/r2/lib/db/_sorts.pyx#L70-L85)

Statistically, it is the lower bound of the [Wilson Score interval](http://en.wikipedia.org/wiki/Binomial_proportion_confidence_interval) at the alpha level based on supplied Z score.

The optional `zScore` parameter can be passed as to the exported `wilsonScore` factory.
The Z score is a statistical value which roughly means how many standard deviations of safety you want, so it maps directly onto the confidence level of the Wilson Score interval.

It will default to `z=1.96` if left out, representing a `95%` confidence level in the lower bound. Otherwise, values through `1.0` (69%), to `3.3` (99.9%) good alternatives.

### 2. Reddit Hot Sort
Based on the difference between ups/downs, and decays with time. Causes hive mind effects in large crowds.

An optional _halflife_ parameter can be passed to the exported `redditHot` factory.
The half-life defaults to 45000 [s]. For info on the effects on this parameter read the original [blog post](https://medium.com/hacking-and-gonzo/how-reddit-ranking-algorithms-work-ef111e33d0d9) about it. See also the canonical [reddit source version](https://github.com/reddit/reddit/blob/bd922104b971a5c6794b199f364a06fdf61359a2/r2/r2/lib/db/_sorts.pyx#L47-L58).

### 3. HackerNews Hot Sort
Based on simply the amount of upvotes, and decays with time. Prone to advertising abuse.

An optional `gravity` parameter (defaulting to `1.8`) can be passed to the exported `hackerHot` factory. For info on the effects of this parameter read the original [blog post](https://medium.com/hacking-and-gonzo/how-hacker-news-ranking-algorithm-works-1d9b0cf2c08d) about it.

## Installation

```bash
$ go get -u github.com/bububa/timedecay
```

## License
MIT-Licensed. See LICENSE file for details.
