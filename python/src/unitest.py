def reverseWord(w):
  return ' '.join(w.split()[::-1])

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
