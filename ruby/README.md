# Usage

`docker build . -t my-app`

`docker run -it my-app /bin/bash`

`root@68ea80f3dcd2:/# cd my-app/my-app/`

`root@6baaebc4b029:/my-app/my-app# bundle install`

`root@6baaebc4b029:/my-app/my-app# bundle exec functions-framework-ruby --target http_example`

or

`root@6baaebc4b029:/my-app/my-app# bundle exec functions-framework-ruby --target event_example`

`docker ps`

`docker exec -it 6baaebc4b029 bin/bash`

Other things tried
```
gem install toys --user-install
export PATH=/Users/kaylanguyen/.local/share/gem/ruby/2.6.0/bin:$PATH
toys system update
gem install bundle
bundle install
toys server http_example
```