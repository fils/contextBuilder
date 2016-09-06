# ContextBuilder

###### An experiment to see if a simple tool can be built to provide small individual data providers a larger context for their data holdings from a defined community of providers that describe their services and capabilities 

#### Intro
The initial approach is a simple crawler that will walk through a whitelist
(later add hypermedia driven walking) of domains.  On these domains we look
for a file provided in a CDF blessed schema listing information about
participating CDF domains and their data holdings.  CDF is initially 
recommending the content of these files to be around services and capabilities
(whatever that means...  I'll try and fill in more on that later).

My personal take on this that I want to be able to be able to build a network
of like minded sites around my domains.  For me, these would be mostly in 
the paleo-geosciences.  If I could pull data types, resource types/classes,
vocabularies (SKOS, OWL mainly) and perhaps some service descriptions like 
Swagger I might be able to do something with these to build a domain "context"
that could provide value to my users.   This is expected to be different 
from a direct resource to resource relation.  Rather this is likely more 
like fuzzy connection based around attributes of resources or overlaps in 
various contexts (vocabularies) of given resources.  

#### Context?
My interest in this project is to actually get some local value out of
what normally just turns into another monolithic master metadata 
catalog (M3C) (trademark pending).   Many of us are starting to extract 
data type and resource data from our data holdings.  Given that, are there 
practices we could use locally that look out into our known community of 
practice and find useful connections?

Given a resource I am presenting to a user, can I find potential relation 
outside my own context (site) that I can present to the user.   Could I 
even present this to a user as a separate view on my resource.   Can I, 
with acceptable effort, answer the question "what else is there that is 
potentially related to what I am looking at now".  GeoLink has been working 
on this as well.  What is proposed here is a complement to that, though 
perhaps someting that can and should be just incorporated into GeoLink.

This question might arise because my resource isn't really what the user 
is after.  If I can then at least provide the benefit of getting them to 
a better resource I have still added value to the users experience.  The 
question might also come up because the user needs more context around a 
resource in order to use it.  

A possible extension  of this is if we combined this with the 
search parameters being used.  So, a user is a looking at a local 
resource he or she arrived at from a search 
term or set of parameters.  Knowing connections to other providers we can 
look for search options that match those terms used locally.  The results of 
remote searches can then be combined with static context data and presented 
to the user as other possible resources of interest.  One can easily imaging 
connections to GeoDeepDive or DataOne as possible options.  Many of these places 
already have API's.  Being able to systematically harvest them into a graph and 
potentially match parameters to them.  All this sounds a lot like what is 
done with Swagger documents and client building so looking there is a good
place to start.  Much of the journey may be done.  Also projects that work on 
query building for example are another set of resources to review.  We don't want
to look anyting like a broker though.  If we end up there just quit.


#### What this is not
I don't have the time or resources to do a full harvester with scheduled 
visits, difference mapping, etc, etc.  Groups like DataOne and others 
are far more capable to address that larger scale task.  I'm also not 
set up to address the type of interfaces people would want on such data.   
This project is to see if a small group can use a lightweight tool to help 
provide context to local data from small defined community of other providers 
exposing information this way.   

#### Code 
The main.go program is a just a copy of the of FetchBot example from 
github.com/PuerkitoBio/fetchbot .
Likely simple.go is also.  From these base codes I hope to modify 
them to address my goals.

The modifications are easy, read the body, look for links to a document
types associated with the CDF document.  Read, validate and parse that document
and route the various elements of that document to functions that then fetch 
further document to build out the context graph. The end results will hopefully 
be a simple RDF graphs that can be used locally to give further context to hosted 
data. 

It might be nice too if we had a means to locate the file(s) of interest via mime type. 
Following the OpenSearch pattern with something like:

```
<link type="application/coalitiondatafacilities+xml" rel="notsurebestrelvalue" title="SiteX" href="http://eaxmple.net/cdf.xml" />
```

Or similar for JSON-LD formats. 

