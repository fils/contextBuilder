# contextBuilder

###### An experiment to see if a simple tool can be built to provide small individual data providers a larger context for their data holdings from a defined community of providers that describe their services and capabilities 


#### For review
* http://jonathanmh.com/web-scraping-golang-goquery/
* https://www.hugopicado.com/2016/09/26/simple-data-processing-pipeline-with-golang.html


#### Foundation thoughts 
The following items list a few thoughts and ideas that form the base for this experiment:

- A site should describe itself.  In the same way a software library exposes APIs or interfaces
 a web site should too.  There exist ways to do some of this now: swagger, void, opensearch, etc
- Arguably things like re3data descriptions documents or others are also part of this approach as
 well.  Providing a metadata foundation for the services and resources.
- A possible  gap many see are data types, units, etc.  Other groups are working on ontologies and
 approaches for data type description and exposure.  These should be followed and used to fill in this gap.
- What we really lack is an index/directory of these descriptions, something like a manifest file.
  One could argue that index.html is the manifest and simple application of HTML practices allow 
  for discovery of these.  OpenSearch does this and something like 

```
<link type="application/earthcubecdf+xml" rel="notSureOfBestRelVal" title="SiteX" href="http://eaxmple.net/cdf.xml" />
```  

could be used to address this.

An alternative approach is something like JSON-LD (ref: http://json-ld.org/)  which would allow 
context and possible connection to hypermedia approaches (ref: http://www.hydra-cg.com/).  This 
JSON-LD file provides a more expressive document that could still be initially linked or 
referenced from index.html.

Either approach works and both have trade offs.  At first I will likely play with both to 
see which one resonates with me.  

The question is can one take all these various documents that already exist and weave them together
 into something useful.  Where useful is to measure them against a set of functions a domain needs.  

Some of my personal use cases are:

- How to connect and exchange data of interest with Neotoma
- this 

Just a questions I have (and there are many)

- do I know how to match my query terms to remote terms
- do I know how to relate data types and measurements from my resources to remote resources
- can I collect resources of interest from a remote site to correlate with my resources?

weaving is aided by semantics..  again, opportunity perhaps.  Does swagger allow for injecting 
semantics into the description?  Is there a JSON-LD version of swagger for example.  
 

#### Goals
The initial approach is a simple crawler that will walk through a whitelist
(later add hypermedia driven walking) of domains.  On these domains we look
for a file provided in a CDF blessed schema listing information about
participating CDF domains and their data holdings.  CDF is initially 
recommending the content of these files to be around services and capabilities
(whatever that means...  I'll try and fill in more on that later).

Later I would hope that I could ditch the crawler aspect and just pull the needed material from a reliable indexer service.

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
perhaps something that can and should be just incorporated into GeoLink.

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
to look anything like a broker though.  If we end up there just quit.


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
<link type="application/earthcubecdf+xml" rel="notSureOfBestRelVal" title="SiteX" href="http://eaxmple.net/cdf.xml" />
``` 

Or similar for JSON-LD formats. 


#### Harvesting and working with JSON-LD and triples

Ref RDF_pro  at http://rdfpro.fbk.eu/
```
Fils:output dfils$ ~/bin/rdfpro/rdfpro @r -w bco-dmoorg.nq wwwopentopographyorg.nq opencoredataorg.nq wwwunavcoorg.nq  @w combined.nq
```

#### Cayley hacking

Gizmo
```
g.V().Has("<http://schema.org/name>").Tag("source").Out().Tag("target").All()

g.V().Has("<http://schema.org/potentialAction>").Tag("source")
.Out("<http://schema.org/potentialAction>")
.Out("http://schema.org/target").Out("http://schema.org/description").Tag("target").All()


repository.Tag("source").Out("<http://schema.org/potentialAction>")
.Out("<http://schema.org/target>").Out("<http://schema.org/description>")
.Tag("target").All()


var hasName =   g.V().Has("<http://schema.org/name>")
var hasAction = g.V().Has("<http://schema.org/potentialAction>")

var repos = hasName.Intersect(hasAction).Out("<http://schema.org/name>").Unique()

var targets = repos.Tag("source").In().Out("<http://schema.org/potentialAction>")
.Out("<http://schema.org/target>").Out("<http://schema.org/description>")
.Tag("target")

var members = repos.Tag("source").In().Out("<http://schema.org/memberOf>")
.Out("<http://schema.org/programName>").Tag("target")

targets.Union(members).All()

```


GraphQL
```
{
  nodes{
    <http://schema.org/name>, <http://schema.org/potentialAction>  {
      <http://schema.org/target> { 
       <http://schema.org/description> 
      }
  	}
  }
}
```

