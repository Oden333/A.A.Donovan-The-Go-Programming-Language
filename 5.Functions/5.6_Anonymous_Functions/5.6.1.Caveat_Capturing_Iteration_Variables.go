package main

// var rmdirs []func()

//* for _, d := range tempDirs() {
//*     dir := d               // NOTE: necessary!
//*     os.MkdirAll(dir, 0755) // creates parent directories too
//*     rmdirs = append(rmdirs, func() {
//*         os.RemoveAll(dir)
//*     })
//* }
//* // ...do some work...
//* for _, rmdir := range rmdirs {
//*     rmdir() // clean up
//* }

//! the for loop introduces a new lexical block in which the variable
//! dir  is  declared.  All  function  values  created  by  this  loop  “capture”  and  share  the
//! same  variable—an  addressable  storage  location,  not  its  value  at  that  particular
//! moment.

//! The value of dir  is  updated  in  successive  iterations,  so  by  the  time  the
//! cleanup functions are called, the dir variable has been updated several times by the
//! now-completed for loop.
