# dns-common

## specific domain

using spec00

example:

spec00-fikcbikcafkcdax-17-thisisthepullzone.domain.com

{spec00}-{node tag}-{pullzone length}-{pullzone name}.domain.com

## gen and check cname setting when using custom cert

example

applyDomain: www.somedomain.com

pullZoneName: pullzonexxx

hostedDomain: mesoncdn.com

gen setting

hostDomainCname: pullzonexxx.mesoncdn.com

challengeRecord: _acme-challenge.www.somedomain.com

challengeTarget: _acme-challenge.www.pullzonexxx.mesoncdn.com

set cname record
1. CNAME  www.somedomain.com => pullzonexxx.mesoncdn.com
2. CNAME  _acme-challenge.www.somedomain.com => _acme-challenge.www.pullzonexxx.mesoncdn.com
