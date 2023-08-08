# What is tricolor garbage collector

// --> GO-lang's GC uses tricolor mark and sweep method to free up the spaces

### Also golang's garbage collector handles the below conditions

 -> Data races
 -> Dangling Pointers
 -> Dereferencing Pointers
 -> Double Free
 -> Memory Leak (in the sense, continous allocation of memory eventually leads to memory overflow )

# working of the GC

```Tricolor Marks and Sweep in Golang

This Tricolor Marks and Sweep algorithm divides objects in the heap into three sets according to the color assigned, such as black, white, and gray. The white set objects are the candidates for garbage collection. The black set objects guarantee to have no reference to any objects in the white set. But note: objects in the white set may have reference to the objects in the black set; this has no effect on the operation of the garbage collector. The objects in the gray set might have reference to the objects in the white set.

Here, Go programmers must understand that the gray set is the intermediary set and no black set object can directly be white or black under any circumstance. They must go through the gray set from where it is decided whether to put in the black or white sets. This means that when the garbage collection cycle begins, all objects are marked as white to begin with, then gray, and, as the algorithm operates, the garbage collector visits the root object and colors them black or white accordingly. The root of the object essentially means the objects that the application can directly access.

Tricolor Scheme in Go

The steps in the Tricolor Scheme can grossly be summed up as follows:

Step 1: Pick an object from the gray set.
Step 2: Gray all the object references and move it to the black set
Step 3: Repeat step 1 and 2 until the gray set is empty.
Step 4: Put all other objects in the black set if they are reachable from the root otherwise put them into the white set.
Step 5: The objects in the white set can be garbage collected.```
