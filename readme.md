# Olivier's Reference Sheet
If I google something more than once, it means that I did not take the proper time to learn everything about this
thing. I am creating this sheet so that I don't fall in the bad habbit of googling everything. I will format this
by categories/Technologies

## Content
1. [R-language](#R-language)
  1. cut {base}

## R-language
### cut {base}
cut divides the range of x into intervals and codes the values in x according to which interval they fall. The leftmost interval corresponds to level one, the next leftmost to level two and so on. logical, indicating if an ‘x[i]’ equal to
    the lowest (or highest, for <code>right = FALSE</code>) ‘breaks’
    value should be included.
    If a <code>labels</code> parameter is specified, its values are used to name
      the factor levels.  If none is specified, the factor level labels are
      constructed as <code>"(b1, b2]"</code>, <code>"(b2, b3]"</code> etc. for
      <code>right = TRUE</code> and as <code>"[b1, b2)"</code>, … if <code>right =
        FALSE</code>.
      In this case, <code>dig.lab</code> indicates the minimum number of digits
      should be used in formatting the numbers <code>b1</code>, <code>b2</code>, ….
      A larger value (up to 12) will be used if needed to distinguish
      between any pair of endpoints: if this fails labels such as
      <code>"Range3"</code> will be used.  Formatting is done by
      <code><a rd-options="" href="/link/formatC?package=base&amp;version=3.6.0" data-mini-rdoc="base::formatC">formatC<a><code>
```
cut(x, …)
# S3 method for default
cut(x, breaks, labels = NULL,
    include.lowest = FALSE, right = TRUE, dig.lab = 3,
    ordered_result = FALSE, …)
```
Example how to cut date into multiple quarter
```
OrderCons$Period<-cut(OrderCons$ORDER_CREATION_DATE ,
                    breaks = ymd(c( "2018-04-28","2018-07-29","2018-10-28",
                                  "2019-01-27","2019-04-28")),
                    labels =c( "Q4FY18","Q1FY19","Q2FY19","Q3FY19"))
```
#### Reference
1. [rdocumentation](https://www.rdocumentation.org/packages/base/versions/3.6.0/topics/cut)
2. [stackoverflow](https://stackoverflow.com/questions/45201474/customize-quarterly-dates-on-r)
