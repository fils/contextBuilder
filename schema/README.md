# Schema notes

My first goal is to simply build a JSON-LD file that points to other records.
My first collection of these being:

* RE3Data files
* VoID (for the semantic folks)
* Swagger (for the API folks)

If I can locate and index these then I have a good 0th order idea of 
what people are offering.  Then perhaps branch out into things like 
data types or measurements.  Once collected the real trick is to map these
data to the local interface/experience to see if there is any value that can
come from them.   If not..  quit, refocus elsewhere.

## RE3

Reference http://service.re3data.org/schema for information about 
the RE3Data schema.  

* Direct link to the Template: http://schema.re3data.org/3-0/re3data-template-V3-0.xml
* Direct link to the Example: http://schema.re3data.org/3-0/re3data-example-V3-0.xml

## JSON-LD

I've focused on the use of JSON-LD for this test for several reasons.

* Ease of parsing into native structures in Go
* Ease of translating to RDF for storage in triple stores.   Alternatively
I could use a JSON store (MongoDB) or map into a KV store (BoltDB) and 
work from there.   Also easy to build full text indexes from JSON-LD

## ToDo

* add into the JSON-LD links to things like git repos and onotosft landing pages.
(need to see if these are machine readable)
* Need to look at what is going on in ESIP and RDA that impact the approach
CDF is trying to take here  (points of contact, invite them to CDF meeting)