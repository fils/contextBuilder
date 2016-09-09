# Connecting Open Core Data and Neotoma

## About
One of the goals of the Open Core Data funded project is to better facilitate
the use of scientific drilling data by "down stream" users.  These focused
domain specific groups offer high value products to a community of practice.

The following is just an initial guess of what a back and forth flow might 
be like between OCD and Neotima.  (need to talk with Simon and make a UML sequence diagram)

* Neotoma asks OCD for datasets related to certain topics defined in 
a OCD vocabulary.  These might be defined be datasets assoicated with 
particular measurements or parameters for example
* a collection of associated datasets and their URI's are sent back to Neotoma along
with something an md5 checksum or other approachs that allow Neotoma to know if they 
have these datasets already or if they have or have not changed.
* also with these datasets would be some form of provenance data for Neotoma 
to use
* Neotoma would keep track of know what datasets (based on UUID) it has and perhaps
also what version (do we need this?) or at least the MD5 to see if a dataset has
changed


Additionally 

* want this to be easy to maintain 
* stateless potentially (easy sync without the need to record stages for example)