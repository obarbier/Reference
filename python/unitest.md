### How test cases is run in python and some pattern I have seen during a code review
I am using the following [example](https://www.practicepython.org/exercise/2014/05/21/15-reverse-word-order.html) to help explain.

Let's say we are asked to Write a program (using functions!) that asks the user for a long string containing multiple words. Print back to the user the same string, except with the words in backwards order. For example, say I type the string:

>  My name is Michele
> > Michele is name My

Let's say after a long time thinking one come up with the following code.
```python
def reverseWord(w):
  return ' '.join(w.split()[::-1])
```

Then a simple test using the unittest library will be as follow.
```python
import unittest
class TestStringMethods(unittest.TestCase):
    def test_reverse(self):
        input = u'My name is Michele'
        output =  reverseWord(input)
        self.assertEqual(output, u'Michele is name My', 'word is not revesed properly')
    def test_is_not_equal(self):
        input =u'My name is Michele'
        self.assertFalse(input == reverseWord(input))
    def test_is_equal(self):
        input =u'My name is Michele'
        test = u'Michele is name My'
        self.assertTrue(test == reverseWord(input))
    def test_assertion(self):
        input =u'My name is Michele'
        print("")
        print("test string is '  {} ' and output is ' {} '".format(input, reverseWord(input)))
        with self.assertRaises(AssertionError):
            # check if test  equal reverse
            # obviously those two string are not equal therefore assertTrue should raise
            # AssertionError type error.  but since we tested for it. the test succeded
            self.assertTrue(input == reverseWord(input))
if __name__ == '__main__':
  unittest.main()
```
A testcase is created by subclassing **unittest.TestCase**. The three individual tests are defined with methods whose names start with the letters `test`. This naming convention informs the test runner about which methods represent tests.
The crux of each test is a call to `assertEqual()` to check for an expected result; `assertTrue()` or` assertFalse()` to verify a condition; or `assertRaises()` to verify that a specific exception gets raised. These methods are used instead of the assert statement so the test runner can accumulate all test results and produce a report.
A list of all assert function

<table border="1" class="docutils">
<colgroup>
<col width="48%">
<col width="34%">
<col width="18%">
</colgroup>
<thead valign="bottom">
<tr class="row-odd"><th class="head">Method</th>
<th class="head">Checks that</th>
<th class="head">New in</th>
</tr>
</thead>
<tbody valign="top">
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertEqual" title="unittest.TestCase.assertEqual"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertEqual(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">==</span> <span class="pre">b</span></code></td>
<td>&nbsp;</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertNotEqual" title="unittest.TestCase.assertNotEqual"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertNotEqual(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">!=</span> <span class="pre">b</span></code></td>
<td>&nbsp;</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertTrue" title="unittest.TestCase.assertTrue"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertTrue(x)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">bool(x)</span> <span class="pre">is</span> <span class="pre">True</span></code></td>
<td>&nbsp;</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertFalse" title="unittest.TestCase.assertFalse"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertFalse(x)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">bool(x)</span> <span class="pre">is</span> <span class="pre">False</span></code></td>
<td>&nbsp;</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertIs" title="unittest.TestCase.assertIs"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertIs(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">is</span> <span class="pre">b</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertIsNot" title="unittest.TestCase.assertIsNot"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertIsNot(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">is</span> <span class="pre">not</span> <span class="pre">b</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertIsNone" title="unittest.TestCase.assertIsNone"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertIsNone(x)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">x</span> <span class="pre">is</span> <span class="pre">None</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertIsNotNone" title="unittest.TestCase.assertIsNotNone"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertIsNotNone(x)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">x</span> <span class="pre">is</span> <span class="pre">not</span> <span class="pre">None</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertIn" title="unittest.TestCase.assertIn"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertIn(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">in</span> <span class="pre">b</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertNotIn" title="unittest.TestCase.assertNotIn"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertNotIn(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">not</span> <span class="pre">in</span> <span class="pre">b</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertIsInstance" title="unittest.TestCase.assertIsInstance"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertIsInstance(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">isinstance(a,</span> <span class="pre">b)</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertNotIsInstance" title="unittest.TestCase.assertNotIsInstance"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertNotIsInstance(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">not</span> <span class="pre">isinstance(a,</span> <span class="pre">b)</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertAlmostEqual" title="unittest.TestCase.assertAlmostEqual"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertAlmostEqual(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">round(a-b,</span> <span class="pre">7)</span> <span class="pre">==</span> <span class="pre">0</span></code></td>
<td>&nbsp;</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertNotAlmostEqual" title="unittest.TestCase.assertNotAlmostEqual"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertNotAlmostEqual(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">round(a-b,</span> <span class="pre">7)</span> <span class="pre">!=</span> <span class="pre">0</span></code></td>
<td>&nbsp;</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertGreater" title="unittest.TestCase.assertGreater"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertGreater(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">&gt;</span> <span class="pre">b</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertGreaterEqual" title="unittest.TestCase.assertGreaterEqual"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertGreaterEqual(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">&gt;=</span> <span class="pre">b</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertLess" title="unittest.TestCase.assertLess"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertLess(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">&lt;</span> <span class="pre">b</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertLessEqual" title="unittest.TestCase.assertLessEqual"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertLessEqual(a,</span> <span class="pre">b)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">a</span> <span class="pre">&lt;=</span> <span class="pre">b</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertRegexpMatches" title="unittest.TestCase.assertRegexpMatches"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertRegexpMatches(s,</span> <span class="pre">r)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">r.search(s)</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertNotRegexpMatches" title="unittest.TestCase.assertNotRegexpMatches"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertNotRegexpMatches(s,</span> <span class="pre">r)</span></code></a></td>
<td><code class="docutils literal notranslate"><span class="pre">not</span> <span class="pre">r.search(s)</span></code></td>
<td>2.7</td>
</tr>
<tr class="row-even"><td><a class="reference internal" href="#unittest.TestCase.assertItemsEqual" title="unittest.TestCase.assertItemsEqual"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertItemsEqual(a,</span> <span class="pre">b)</span></code></a></td>
<td>sorted(a) == sorted(b) and
works with unhashable objs</td>
<td>2.7</td>
</tr>
<tr class="row-odd"><td><a class="reference internal" href="#unittest.TestCase.assertDictContainsSubset" title="unittest.TestCase.assertDictContainsSubset"><code class="xref py py-meth docutils literal notranslate"><span class="pre">assertDictContainsSubset(a,</span> <span class="pre">b)</span></code></a></td>
<td>all the key/value pairs
in <em>a</em> exist in <em>b</em></td>
<td>2.7</td>
</tr>

</tbody>
</table>

## teardown and setup
The test case and test fixture concepts are supported through the **TestCase** and **FunctionTestCase** classes; the former should be used when creating new tests, and the latter can be used when integrating existing test code with a unittest-driven framework. When building test fixtures using TestCase, the `setUp()` and` tearDown()` methods can be overridden to provide initialization and cleanup for the fixture. With FunctionTestCase, existing functions can be passed to the constructor for these purposes. When the test is run, the fixture initialization is run first; if it succeeds, the cleanup method is run after the test has been executed, regardless of the outcome of the test. Each instance of the TestCase will only be used to run a single test method, so a new fixture is created for each test.
### TestCase Example
```python
import unittest

class WidgetTestCase(unittest.TestCase):
    def setUp(self):
        self.widget = Widget('The widget')

    def tearDown(self):
        self.widget.dispose()
        self.widget = None

    def test_default_size(self):
        self.assertEqual(self.widget.size(), (50,50),
                         'incorrect default size')

    def test_resize(self):
        self.widget.resize(100,150)
        self.assertEqual(self.widget.size(), (100,150),
                         'wrong size after resize')
```

### FunctionTestCase example
```python
def testSomething():
    something = makeSomething()
    assert something.name is not None
    # ...
testcase = unittest.FunctionTestCase(testSomething,setUp=makeSomethingDB,tearDown=deleteSomethingDB)
```


## reference
1. [pythonDoc](https://docs.python.org/2/library/unittest.html)
