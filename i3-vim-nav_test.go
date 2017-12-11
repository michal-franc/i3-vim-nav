package main

import "testing"

func TestVimRegexp(t *testing.T) {

  testCases := map[string]bool{
    "vim": true,
    "i3-vim-nav.go (~/tools/i3-vim-nav) - vim": true,
    "i3-vim-nav.go": false,
    "VIM": true,
    " vim": true,
    "vim ": true,
    " vim ": true,
  }

  for key, val := range testCases {
  actual := isItInVimContext(key)
  expected := val
  if actual != expected {
    t.Errorf("For %v expected %v but got %v", key, expected, actual)
    }
  }
}
