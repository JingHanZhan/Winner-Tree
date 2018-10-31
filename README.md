# Winner-Tree
Use Winner Tree to complete the part of merge of external sort.

# Why we need external sort?
When we try to sort some file, we can load all of them and sort if memory is enough. It's called internal sort. But if memory is not enough to load all the data, we need to sort with the disk.   

# Ideas
If the number of files is 2^n, they can compose full binary tree. It is ok to merge all of them at same time. But when the number of files that we want to merge is not 2^n, It may be some problem there. 


I try to use three method to complete my Winner Tree merge. 
