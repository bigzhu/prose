package prose

import (
	"encoding/json"
	"log"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var longArticle = `
000.2200 Hi.
001.2200 I'm Vanessa from SpeakEnglishWithVanessa.com.
004.0600 Today we're going to talk about kitchens.
008.10900 Let's go.
0014.600 No matter where you live around the world,
all houses have something in common.
0019.02900 They have a kitchen.
0020.5800 Maybe your kitchen just has one pan and one
burner or maybe you have a huge kitchen with
0026.2300 pots and pans passed down from your grandmother.
0029.500 Whatever the case may be, there are a lot
of varieties of kitchens around the world
0033.6900 and there's a lot of varieties of kitchens
in the US.
0036.6700 So I can't explain all US kitchens, but I
can explain my kitchen.
0041.6600 So that's what we're going to do.
0042.9100 I'm going to take you on a little tour around
my kitchen and explain all of the items that
0047.5100 I see, all of the appliances.
0049.8600 Maybe there's some things that are different
in your country than what are in my house,
0054.5100 but it's a good way to expand your vocabulary.
0056.9700 Because if you spend any time in your kitchen,
it's a great chance to just talk about what
0062.62100 you're doing in English.
0064.01900 I'm cooking an egg in my frying pan.
0066.6700 I'm using a spatula.
0068.6400 Talk about what you're doing.
0070.000 And I hope that today's lesson will give you
the vocabulary to be able to do that.
0073.6300 If you're interested in learning other household
words not just the kitchen, you can check
0077.6100 out this video that I made, 150 Household
Items, and expressions, and verbs, in my previous
0085.3900 apartment.
0086.3900 Now we're in our new house, which is great.
0088.8400 Or if you want to learn more about cleaning
expressions, because that's something that
0092.4100 we all have to do and you want to describe
it in English, I made a video at my mother-in-law's
0097.1200 house using a lot of cleaning expressions.
0099.9400 I think there was 70 or 80 of them.
00102.61900 You can watch that video up here to expand
your vocabulary.
00106.3500 Let's start by talking about some common kitchen
appliances.
00109.4600 One of the most common appliances is the fridge.
00114.7400 Or you can say refrigerator, but this is a
pretty big word so we often cut it down and
00120.6700 just say fridge.
00121.8500 Good news for you.
00124.6300 In the US we have giant fridges.
00127.47900 I've never been in any other country that
had fridges as big as in the US.
00132.7500 Maybe that's kind of typical.
00133.800 We like to make things super big.
00137.0900 But on the top in my fridge we have the fridge
part and on the bottom is the freezer.
00143.1900 We'll take a look at that in just a second.
00145.31900 It's pretty typical that in the fridge you
have different shelves.
00151.5600 It's a little bit bright on my camera.
00153.0200 I'm sorry.
00154.0200 And then at the bottom we call these vegetable
crispers, just a crisper.
00161.9100 We don't use that that often but it's the
crisper bin.
00164.5100 It's somewhere where you'll put your fresh
vegetables, fresh fruit to keep them crisp,
00170.5700 and that is fresh, exactly the way you want
them.
00173.7900 On the bottom of my fridge is the freezer.
00178.1400 Sometimes the freezer will be kind of split
in half and will be on one side or the other
00183.2600 but mine is a top and bottom type of deal.
00187.0200 So on the bottom is the freezer, which is
where we keep some leftover soup, some frozen
00195.34900 vegetables, frozen fruit, ice, ice cream.
00199.6200 Anything that you want to keep super cold
you put in the freezer.
00202.7600 The second most common kitchen appliance is
a stove and an oven.
00209.0100 In the US we call the top of this the stove
or you might say the stove top.
00215.6900 That's where you cook things with pots and
pans.
00220.8800 Inside the oven is where you bake things.
00224.0800 So usually there's a door on the oven and
inside here there are some racks.
00232.4200 That's what we call these things.
00233.7600 It's a little bit hot in there so I'm not
going to grab them, but there are rack.
00238.0800 You might see a recipe that says, "Put this
on the bottom rack," or, "Put this on the
00242.900 top rack when you're baking it," and that's
referring to those kind of metal shelves inside
00248.900 your oven.
00250.0100 You might here some Americans calling all
of this, this whole device, the oven or call
00257.07900 the whole thing the stove instead of separating
the top is the stove, inside is the oven.
00264.9100 Sometimes we just kind of use one of those
words to talk about the whole device, so you
00269.5700 might hear that in movies and TV shows.
00272.0900 It's technically the stove and the oven, but
you might kind of interchange those sometimes.
00279.1100 Some technical words about the stove top is
we have the knobs.
00283.4400 There are knobs to turn on the stove.
00288.5300 And when you are baking something in here,
you need to set the temperature.
00294.2200 In the US we use Fahrenheit.
00297.1800 So when I turn on the oven and I click bake,
the first temperature is 350 degrees, and
00306.2700 that is in Fahrenheit.
00310.1100 If you're cooking in the US, make sure that
you know if your recipe is American, if your
00315.6700 recipe is not from the US, if your stove is
using Fahrenheit.
00319.63900 I've never seen a stove in the US that didn't
use Fahrenheit, but I know when I've baked
00324.6200 things in other countries I had to convert
everything and it was a little complicated.
00329.0600 So don't mess up your food.
00330.69900 Just check on the temperature.
00332.8300 Check on what you're using for the recipe.
00335.300 That's all really essential.
00337.2900 One thing also, above the stove is the microwave.
00341.9900 This is a really typical placement for a microwave
in an American kitchen, maybe in your kitchen,
00346.9900 too.
00347.9900 The reason why they put a microwave above
the stove is it's a convenient location.
00354.9900 But also, listen to this.
00361.7100 What do you think that whirring sound is?
00366.0600 It's a fan.
00367.0600 So underneath here there's a fan.
00369.1400 I have a light on as well.
00370.7800 You can see that light.
00374.69900 And that fan is going to help to suck up all
of the smells, and steam, and anything that
00382.1700 you're cooking here.
00383.5800 It's going to help that not spread out throughout
the rest of the house.
00387.0900 So in my old apartment, we didn't have a microwave
and there was no fan above our stove.
00393.300 So whenever we cooked something, our entire
house got so hot and it smelled so strongly
00400.8800 like whatever we were cooking.
00402.63900 Sometimes that's good, sometimes that's bad.
00404.8100 But, it's really nice to have a fan so that
all the smells and all the steam can get sucked
00410.9600 up and the rest of your house isn't completely
humid from your cooking.
00415.38900 The last common item in an American kitchen
is a dishwasher.
00420.16900 I love the dishwasher so much.
00423.8700 In our old apartment, we didn't have a dishwasher
and we bought this tabletop dishwasher.
00428.7300 It was so amazing.
00430.7300 If you live in a place where you're always
hand washing dishes, I respect you a lot because
00437.0500 that is tough work, especially when you're
cooking a lot.
00440.500 We cook a lot at home.
00442.0800 When you have kids, it is not easy to spend
time doing the dishes by hand.
00447.700 So you can do the dishes by hand or you could
use a dishwasher.
00451.47900 Yes, thank you, dishwasher for existing.
00454.25900 So inside the dishwasher there are racks.
00457.7200 You might recognize that word from the oven.
00461.7400 Inside the oven there are oven racks and there
are racks in here as well.
00465.8300 Something you might notice is on some dishes,
on the bottom ... This one doesn't say it.
00472.4800 It might just say dishwasher safe, but on
some dishes they might say top rack safe only.
00479.4900 For example, we have these cool little bamboo
bowls that we use for Theo, our two-year-old,
00488.9600 and we use them, too.
00491.13900 There's some plates to go along with it.
00493.7400 These are our dirty dishes.
00495.6900 But when I bought these, it said top rack
safe only.
00501.1800 That's because there's also a bottom rack.
00505.8700 And this rack usually gets a little bit hotter,
and it's going to have a little more pressure.
00512.27900 It's just going to be a little bit more intense
for whatever you're washing.
00517.75900 So some things might say top rack safe only.
00520.78900 And if I'm not sure, sometimes I just put
it on the top rack just to be careful with
00525.17900 it.
00526.17900 Next, I'd like to just simply go around the
kitchen and talk about what's out on the counter,
00531.51900 the things that are visible, and then we're
going to go into the cupboards and into to
00537.6700 drawers and discuss what's in there.
00539.30900 It could be appliances, could be things you
eat with.
00541.9900 It could be food.
00544.3800 Let's start on the counter.
00545.3800 So this is the counter.
00548.5500 Simple word.
00549.5500 Or you might call it the countertop.
00550.82900 This is where we prepare food.
00552.9800 It's where we set out food.
00555.4900 This countertop is kind of unusual.
00557.58900 Usually you'll see laminate counters, which
we'll take a look at over here.
00561.24900 You might be able to see a little bit right
here.
00563.6900 Kind of like a plastic-type surface.
00567.27900 Or you might see granite countertops.
00569.7900 That's stone.
00571.33900 It's really nice.
00572.600 But this is a wood block.
00577.0800 The family who lived in this house before
us, the woman was a professional baker so
00582.7600 she had this wooden block to replace the old
countertop and now we get to use it, which
00589.87900 is pretty cool.
00591.25900 So it's kind of unusual to see a big wooden
block like that but usually you'll have a
00595.1100 big counter where you can prepare food.
00597.01900 And this is one is especially big, which is
wonderful, because in our old apartment we
00600.77900 had a small counter space for preparing food
and now we can prepare everything, which is
00605.9700 great.
00606.9700 The next thing that you'll see on this counter
is a fruit stand.
00612.700 We have some fruit hanging out here and some
fruit bowls.
00616.1600 Then, Dan's precious coffee maker.
00618.69900 If you've been watching my channel for a while,
you know that I don't drink coffee but Dan
00625.01900 loves it.
00626.01900 So we'll get to what I have over here.
00627.81900 So we have the coffee maker.
00631.2900 Your coffee maker might look a little more
traditional than this, but this is Dan's special
00636.6700 thing.
00638.33900 Then we have the paper towel stand.
00642.26900 Sometimes this paper towel stand might be
underneath a counter and it kind of attaches
00648.58900 to the bottom of the counter and you can roll
it out underneath a shelf, something like
00653.1700 that, but ours just is a standalone paper
towel rack.
00659.0700 And at the sink we got a couple other things,
so let's go take a look.
00662.52900 This is our sink where we wash things, where
we fill up water.
00667.2900 All of those great things happen.
00668.8700 At most sinks you'll see dish soap.
00673.64900 It's quite clear.
00675.2900 Liquid dish soap.
00676.2900 It's easy to understand.
00677.83900 This is what you're going to use when you're
hand washing dishes.
00682.2300 You do not, do not, want to put dish soap
into your dishwasher.
00690.200 If you put dish soap, and it usually says
something about ... Yes, dish soap for hand
00698.04900 washing dishes that uses biodegradable, bio-based
formula, blah, blah, blah.
00702.1100 So it says hand washing.
00704.57900 That's your key.
00706.1100 If I put this in my dishwasher, our house
would be covered in bubbles.
00711.54900 So please don't do this unless you want to
have some fun experiment.
00716.4400 There is specific soap you put in the dishwasher.
00721.2300 It might be a little pod like this or it's
... Usually we buy just powder that you put
00729.63900 in the dishwasher.
00730.63900 It's a little box that has dish washing soap.
00735.39900 So make sure that you use dishwasher soap.
00739.38900 It's for washing your dishes in the dishwasher
in the right place, and you use this in the
00743.9700 right place.
00745.55900 Little friendly word of warning.
00746.9200 So usually at the counter, when you're washing
dishes, you'll find some kind of towel.
00751.92900 We don't use sponges, but you might find a
sponge.
00755.4100 This is like a reusable, cool little microfiber
towel.
00759.63900 We have a bunch of these, and that's what
we use to hand wash dishes.
00763.4700 You might also find ... This is called steel
wool.
00767.69900 It's just a little piece of metal mushed up
together.
00773.76900 You can use it to kind of scrape some hard
spots on pans or pots.
00778.3100 You'll also see a faucet.
00780.44900 This is the faucet.
00782.05900 And hot water, cold water.
00786.20900 We're so lucky we're living in the 21st century
we can just have hot and cold water whenever
00790.4700 we want.
00791.4700 It's amazing.
00792.4700 Our sink also has this cool little thing.
00795.7600 I wonder if this exists in your country.
00798.33900 I'm going to turn on the water, and then I'm
going to turn this.
00804.27900 Whoa.
00805.27900 Isn't that cool?
00807.38900 I'm amazed.
00808.9500 This is just like a little sprayer so you
can spray the sink out or you can spray a
00816.80900 dish.
00818.52900 I once played a terrible prank on my dad for
April Fools' Day because if ... This has a
00827.62900 little button on the side, and I put a rubber
band around it so it was always pressed.
00836.27900 When you turn on the water, usually it comes
out here, but if this button is pressed, it
00845.19900 will come out of the sprayer automatically.
00850.33900 So I had a rubber band around it.
00852.35900 And in the morning, he went to the sink and
he turned on the water.
00856.600 You can imagine what happened.
00858.24900 It didn't come out here.
00859.57900 It came out here and went ... and sprayed
him.
00864.08900 It was a good prank.
00865.31900 And it was April Fools' Day, so it's all right.
00867.4700 But, I thought that was a pretty clever moment
when I was in high school, so watch out for
00872.76900 rubber bands around your sprayer.
00876.55900 There's also hand soap.
00877.82900 This is for your hands.
00879.6800 That's quite obvious.
00881.39900 Usually we use liquid hand soap.
00883.4200 You might see people use a bar of soap, but
that's pretty unusual.
00886.1800 I am curious in your country if you use bar
soap or liquid soap like that.
00892.70900 Let's continue on to the side and see some
other things that are on the counter.
00896.55900 You might be wondering what this giant metal
tub is.
00900.80900 Well, this not common in most American households,
but I'm a water snob and I love clean, pure
00912.70900 water.
00913.70900 In the US, it's totally fine to drink tap
water.
00917.37900 That's water that comes from the faucet.
00919.77900 It's not a problem I'd say 98% of places in
the US.
00926.82900 You can drink the tap water.
00928.00900 And here, you can drink the tap water.
00929.89900 It's completely fine.
00931.88900 But I think it tastes very strongly like chlorine
and just kind of chemicals.
00938.31900 So we have this water filter, and we use this
glass jar, take the tap water, and pour in
00946.19900 here.
00947.19900 It has some huge filters.
00950.1900 This is called like a gravitational long filter
device.
00955.50900 I forget the technical term for it.
00957.39900 But, it doesn't filter it automatically.
00959.43900 It takes maybe two hours to filter so we're
always pouring water in there.
00964.42900 It stores the water down here, and it tastes
amazing.
00967.4100 So if you ever come to my house, drink some
water.
00971.64900 All right.
00972.6700 Another part of this counter is ... This is
kind of a fancy, fancy-shmancy device, but
00978.4100 I got for Dan one year for Christmas because
he likes to drink sparkling water.
00983.72900 I felt like it was a little bit wasteful to
buy cans of sparkling water so I bought this
00992.60900 soda stream device.
00994.26900 It's pretty nifty.
00995.62900 So there's a bottle, a specific bottle, that
goes with this machine.
00999.1900 And you fill the bottle with water.
001001.100 You put it in here, and it just carbonates
the water, so it makes it sparkly.
001007.6600 You can do it as much as you want.
001008.97900 You can use your own water, and it's kind
of an endless thing.
001013.44900 Then when your carbonation device is empty,
you can just return it and they recycle it.
001020.12900 It's a really cool system.
001021.12900 Anyway, if you drink a lot of sparkling water,
check it out.
001024.57900 Well, the next thing is a toaster.
001027.3900 You might have this in your house or maybe
you don't.
001029.6400 In my other video, 150 Household Items, we
had a toaster oven.
001034.8700 That's with a door, and you have to lay everything
flat.
001038.16900 You can also use it like an oven.
001039.68900 It's kind of like a mini oven.
001041.95900 But In this house we decided to get a toaster,
just a regular toaster, that has slots at
001047.75900 the top.
001048.75900 Then we have Dan's coffee grinder, more coffee
devices.
001053.0900 And here is my section.
001054.7900 This is the tea section.
001056.9700 I have a lovely teapot, some tea strainers,
a couple things of tea right here.
001063.9600 I have another drawer that's got lots of fun
tea.
001067.0700 If you drink tea in your country, I would
love to try tea from your country.
001073.66900 That's always the thing that I love to try.
001075.8200 I love to try lots of stuff when I travel.
001077.55900 But whenever I travel, I love to try the kind
of traditional food and also the kinds of
001082.65900 tea that people drink around the world.
001085.0100 So I'm always curious what kind of things
that other people drink.
001088.3200 My PO Box is in the description if you want
to send me any.
001091.90900 All right.
001092.90900 We are back at the stove top.
001094.5700 But on the stove top there is something I
didn't mention, and that's a tea kettle for
001098.9900 heating water.
001100.5700 Pretty clear it's for my tea or for something
else if you want to just heat up some water.
001105.7400 Got a cool little thermometer on the top so
that you can see how hot it is because you
001111.7400 don't want to overheat your green tea.
001114.1400 It must be the perfect temperature.
001116.4600 Maybe some people don't care, but I do.
001118.6300 And finally, before we go into the cupboards
and the drawers, I'd like to end at this piece
001124.30900 of the counter which has some important elements
for cooking.
001128.26900 We have some cutting boards.
001130.66900 These are wooden cutting boards that we have
here, and this kind of laminate wood cutting
001137.2400 board here, and some oven mitts.
001140.6700 So this is a flat oven mitt and this is an
actual mitt, like a mitten, that you can use
001148.54900 to protect your hands.
001150.6300 On this little tray ... Not everybody has
this.
001152.6100 This is just something that we do for convenience.
001155.04900 We have some common items that we use when
we're cooking, like olive oil, or honey, or
001159.67900 salt and paper, or my prenatal vitamin for
our baby, some chopsticks, some knives, and
001167.7400 some kitchen scissors.
001169.8800 I'm curious if you use scissors in the kitchen.
001172.7200 We use scissors a lot, but I feel like growing
up my parents didn't use scissors that much
001179.7700 in the kitchen.
001180.7700 So maybe it's kind of a more modern thing.
001183.400 I'm not sure.
001184.400 Do you use scissors in the kitchen?
001186.3800 Next, let's take a look at some of the dry
goods, some things in cupboards, some different
001192.4900 appliances, and things that are put away.
001194.6200 These shelves are quite high, as you can see.
001197.75900 But up here we have some dry goods.
001200.7500 Dry goods are any kind of food that doesn't
need to be refrigerated.
001204.50900 So for us, that's pasta, or couscous, or quinoa,
the types of things that don't need to be
001211.42900 refrigerated.
001212.42900 We have some alcohol.
001213.5200 I was actually using that to make some apple
cider spiked with rum, not for me, unfortunately,
001223.4400 because I'm pregnant, but for our friends
for our pumpkin party we had.
001228.51900 Here is a little bit messy, but it's all right.
001231.17900 We have cookbooks.
001232.3900 A lot of kitchens have cookbooks.
001233.89900 You probably do, too.
001234.89900 If you have any favorite cookbooks, please
recommend them.
001237.8900 I always love browsing through cookbooks.
001242.1300 And at the end of every year for Christmas
I usually buy a new cookbook.
001247.3300 We have just a couple here, but it's always
nice to start the year with new recipes and
001253.7100 new ideas.
001254.7100 So if you have any recommendations, let me
know.
001256.55900 I have some cookbooks.
001257.55900 This is my personal cookbook.
001259.6200 Not really my recipes but these are just recipes
that I have written down in a little notebook.
001265.2600 We also have some other ... These are mainly
knickknacks, just kind of for display.
001269.3700 It's not really for the kitchen.
001270.91900 But here we have other dry goods.
001273.1100 These dry goods are kind of our snacks goods.
001277.5900 So we try to have healthy snacks, especially
with a two-year-old who always wants snacks.
001283.100 I don't want to give him something unhealthy
all the time so we always try to have a lot
001287.52900 of nuts available.
001289.300 So in these different jars ... These are usually
called mason jars.
001295.92900 It's the brand of jar but we use it as just
a name of it.
001300.5700 It's a mason jar or you could say it's a glass
jar.
001303.55900 It has this type of lid that's kind of two-piece.
001307.05900 It's got two pieces to it.
001310.1900 So we usually store different peanuts, cashews,
almonds, pistachios, Brazil nuts, popcorn
001320.9600 seeds, or sunflower seeds, or raisins, anything
that's kind of a healthy snack we try to put
001329.0900 on this shelf.
001330.0900 All right.
001331.0900 Let's go back to the kitchen, which is just
right here, and we're going to take a look
001334.55900 inside some of the cupboards to look at some
other appliances and things that might be
001338.200 hiding.
001339.40900 Pretty much every kitchen has something that
you eat on, right?
001342.7200 So we have some small plates.
001345.8300 There's not really a fancy term for these.
001347.9700 You might have really small plates that are
called saucers.
001351.8500 We don't have these, but you can use it to
put like a teacup on.
001356.90900 And it usually has a little rim on the inside,
some place where you can safely set your teacup.
001364.04900 This is not a saucer.
001367.0900 This is just a small plate for small portions
of food.
001370.8600 Then we also have some big plates, some larger
plates.
001377.1100 We do have bowls, but, as you saw, they are
dirty.
001380.8300 All of our bowls are dirty.
001381.9500 We usually go through bowls really quickly
because, especially now, it's a little bit
001386.3600 colder outside so we eat a lot of soup in
our house and we finish all of our bowls pretty
001392.77900 quickly.
001393.77900 So we do have bowls.
001395.1700 I'm sure you know what bowls are, though,
so no problem.
001397.9400 On this shelf as well we have some mugs, some
more mugs in here.
001402.2400 I love mugs.
001404.4600 And some glasses or just some cups.
001408.17900 You can call this ... It is glass.
001411.1300 It is made of glass, but you can also call
it just a cup.
001416.7200 There's some smaller cups.
001418.3200 There's some beer glasses.
001420.4700 This is the right size for drinking beer according
to Dan.
001423.7500 I don't know.
001424.7500 I'm not a big fan of beer.
001426.4500 But we also have up here a strainer.
001431.54900 Sorry.
001432.89900 I'm a little bit low here because these shelves
are really high.
001438.600 This is a strainer.
001439.600 We put this in the sink to wash our fruit
and vegetables.
001443.8200 Super handy.
001444.8200 You probably have something like that in your
kitchen.
001446.4800 And we just have some other coffee things,
like a hand grinder.
001453.4700 Dan loves grinding coffee by hand, or a pour
over for coffee.
001458.55900 This is kind of the miscellaneous coffee supplies.
001461.62900 There's a French press up there, other little
devices.
001465.63900 All right.
001466.63900 Let's go to this shelf over here.
001469.5900 Okay, I'm just going to sit up on the counter
because it's a little bit awkward to film
001474.200 a high spot and I'm down here and I'm showing
this.
001477.9800 So, I'm siting on the counter.
001479.53900 No judging.
001480.53900 All right.
001481.53900 Up here we have our wine glasses.
001484.7300 Your wine glasses might have a stem.
001489.1200 A stem is that part on the bottom.
001491.3600 A stem is the word that you use with a flower.
001495.55900 There's a flower, and then under the flower
is a stem and it looks like that.
001499.3600 But, these wine glasses do not have a stem.
001501.9700 We have some little fancy glasses, too.
001505.92900 And on this shelf is generally just knickknacks.
001509.5700 This is a point of contention in our relationship.
001514.8600 Whenever Dan has a really cool beer glass,
he wants to keep it because it looks cool.
001520.8300 But I feel like when you display too many
beer glasses it kind of looks bad.
001529.7300 Kind of looks like you drink too much alcohol.
001532.76900 And Dan doesn't drink that much alcohol, but
when you display all of it, it's not that
001538.2400 great, I think.
001539.63900 Anyway, so I'm always secretly recycling them.
001543.12900 Shh.
001544.12900 Don't tell him.
001545.12900 Actually, he knows.
001546.12900 He knows.
001547.12900 Sometimes he'll look up and say, "Hey, where'd
that go?
001549.42900 Oh.
001550.42900 Okay.
001551.42900 I guess there were too many."
001552.42900 So these are some glasses that Dan thinks
are super cool.
001555.55900 All right.
001556.55900 Now we're going to go around the kitchen drawer
by drawer.
001558.4700 I'm going to pull some things out.
001559.8600 You're going to see inside the drawers of
kitchen.
001563.01900 Not so clean, but maybe yours is not so clean
either.
001566.87900 All right.
001567.87900 Inside this drawer there are all of our lids
for our pots and pans.
001571.6200 We haven't looked at our pots and pans yet.
001574.05900 We'll take a look at those in a moment.
001575.05900 They're kind of hanging up in an awkward spot
to film.
001577.9600 But we have like a lid rack so that you can
hold all of the lids.
001583.53900 It's so convenient.
001584.5900 I really like having that.
001586.7100 Down here we have a lot of our baking things.
001589.65900 So there's cookie sheets or baking sheets.
001596.26900 You can use these for baking cookies, for
baking french fries, whatever you want to
001602.1200 cook, baking vegetables.
001603.50900 We also have a ... This is technically called
a cooling rack.
001608.7200 So you can bake cookies and set them here.
001611.91900 And because there is some space, it will cool
the cookies and they won't get too moist.
001617.6900 It'll cool evenly.
001619.00900 But I also use this for baking.
001621.1500 It fits on the sheet perfectly.
001624.03900 I use it when maybe we're baking something
that has a lot of grease, like chicken wings,
001630.01900 something that has a lot of grease and I don't
want my food cooking in the grease.
001633.0800 I'll put that up here and then the grease
will fall down underneath onto the pan.
001638.3400 And it's a little bit healthier and also tastes
good for some things to cook it like that.
001642.99900 Some other baking items that are under here
are different pans for baking different items.
001650.26900 So this is a bread pan but we also have different
casserole dishes.
001656.38900 This is a casserole dish.
001658.24900 It's pretty common in the US to eat casseroles.
001662.5800 I don't know if I've ever actually made one
myself.
001666.4500 Maybe that's more of my parent's generation
making a casserole.
001669.12900 I don't know.
001670.12900 Maybe in colder places in the US they make
casseroles.
001673.00900 It's basically just throwing everything in,
maybe some potatoes, maybe some paste, some
001679.1100 cheese, some broccoli, throwing maybe some
meat and you just kind of bake it together.
001686.0900 It's kind of like a strange lasagna if you
can imagine this kind of idea.
001692.0900 It's kind of an easy thing to make.
001693.9800 So this is a casserole dish, but I use it
for other things, maybe for brownies or for
001700.79900 baking vegetables.
001702.3200 I use it for a lot of different things.
001704.15900 This one is a glass baking dish.
001707.5200 Some people might actually call it by the
brand name.
001710.000 The brand name is Pyrex.
001712.88900 That's a really common brand name for baking
items.
001717.24900 Sometimes we use that to describe it.
001719.3400 So instead of saying the glass dish we might
say the Pyrex dish.
001725.7300 Not everyone's going to use that.
001726.7700 They might just say the glass dish but you
might occasionally hear that, and that's just
001730.27900 the brand name being used for the item.
001732.6600 Even if it's not that brand name but it's
that same idea, it's a glass dish, you might
001737.49900 hear that used.
001738.7300 Something else that I dug out from the back
of this shelf are two silicone muffin tins.
001747.27900 Now, you might have ... It's more likely to
have metal muffin tins, but I like to make
001755.00900 muffins and it's really annoying to get them
out of the metal muffin tin, and this you
001763.0400 can use less oil on it.
001765.89900 It just comes out so easily.
001768.2200 It's really great and it lasts for a long
time.
001771.30900 It's very easy to clean.
001772.99900 So I have these silicone muffin tins.
001775.8900 You might use that or just something to make
muffins with.
001780.500 Also in this drawer is a strainer.
001783.29900 Or, this is not a strainer.
001785.37900 This is a steamer, a steamer.
001787.2300 We already saw the strainer, which was up
in the cupboard, and we use that to wash the
001792.2200 vegetables.
001793.2200 This steamer is going to fold up.
001798.54900 It's supposed to fold up.
001801.2500 You can put that in a pan or a pot of steaming
water and you can steam broccoli, you can
001807.04900 steam green beans, you can steam vegetables
in this really easily, steam shrimp.
001813.02900 I use that for all of these things.
001814.54900 I use this steamer basket.
001817.1900 I think these steamer baskets, maybe they're
pretty common around the world, but this folding
001822.000 metal type, I've had a couple of them and
they always fall apart.
001829.1400 The little leaves on them always fall off,
even if it's a pretty good, brand name steamer
001838.2100 basket.
001839.2100 So I'm curious in your country if you use
this type of thing and if you have better
001843.3400 quality ones.
001844.9200 It's cool that it can fold up, but sometimes
the little leaves get stuck.
001849.8500 Okay, that worked pretty well.
001852.4900 But maybe it's just because I have a two-year-old.
001855.4700 He likes to pull them off and then you find
them around the house.
001859.17900 Not very handy when you have a two-year-old.
001860.5700 Next, lets move on from this cabinet and go
to these drawers.
001866.05900 A cabinet is going to open like this, like
a door, and a drawer, it's a tricky word to
001873.2200 say, is going to open like this.
001875.16900 Well, let's start at the top.
001877.1900 At the top we have silverware.
001881.9600 Some people call this flatware.
001884.5800 It's kind of a restaurant word.
001889.3200 When we're talking about forks, spoons, butter
knives, chopsticks, steak knives, children's
001902.9400 stuff, when we're talking about these items
that you put on the table when you eat, the
001910.30900 most common word is silverware.
001912.9800 It doesn't need to be actually made of silver.
001916.1900 Usually they're not because that takes a lot
of work to keep up.
001920.2200 But we usually call that silverware.
001921.6100 At a restaurant, you might hear the restaurant
staff say flatware.
001927.15900 Or maybe if you work in a really fancy restaurant
or ... Yeah, I don't think I've ever heard
001933.29900 a family call it flatware at home.
001936.88900 I just remember in my home ec class ... Home
ec is home economics where you learn about
001943.9500 baking, and sewing, and all those kinds of
things in high school.
001948.200 In my home ec class, I remember our our home
ec teacher said, "In this class you will not
001953.12900 call it silverware, you will call it flatware,"
and everyone kind of laughed a little bit
001958.26900 like, "Who says that?
001959.8200 That's so fancy."
001961.4900 But in her class we had to call it flatware,
but I've never used that term since.
001965.66900 So if you're from the US or you have American
friends, I'm curious if you've ever heard
001971.8600 flatware in that kind of casual situation.
001974.3600 All right.
001975.3600 In the next drawer we have some other helpful
items for baking and cooking.
001979.87900 This is kind of our random drawer.
001984.2100 Everyone has stuff like that.
001985.7300 Here we have teaspoons and tablespoons.
001988.8500 So as I mentioned, ovens use ... Maybe I didn't
mention this but ovens use Fahrenheit in the
001995.0700 US.
001996.0700 And instead of using milliliters, and liters,
and these kind of easy measurements, we decided
002001.8500 to do things the hard way in the US and use
tablespoons, which is this measurement or
002011.76900 this one here, this big circle, or teaspoons,
which are these.
002020.02900 There's various increments, half a teaspoon,
a forth of a teaspoon, an eighth of a teaspoon.
002026.62900 Yeah.
002027.62900 If you're using spices, or salt, or you're
baking, that's really essential.
002033.91900 If you're cooking something, like you're making
a soup, I just throw stuff in.
002039.8600 I don't usually measure it, but that's a personal
preference.
002042.92900 We also have some tongs.
002046.27900 Great for picking up out things when you're
cooking them, like corn, when you're boiling
002051.60900 corn.
002052.60900 Great to have this.
002053.60900 We also have a thermometer.
002056.91900 This is a meat thermometer.
002058.08900 I guess you could use it for water, too, or
for soup, other things, but we usually use
002064.37900 this for meat.
002065.800 On the back of the thermometer it says C/F,
and that's Celsius of Fahrenheit.
002073.11900 So if you buy a meat thermometer in the US,
make sure that you have the right button clicked.
002080.2700 And if you're thinking about the temperature,
what temperature should pork be in Celsius,
002088.30900 make sure that you're showing the Celsius
number.
002091.1900 I've thought about that a lot because sometimes
the wrong one is clicked and in my head I'm
002096.200 thinking the Fahrenheit number but it's showing
the Celsius number and it's just mixed up.
002101.0100 So make sure that don't get that wrong, especially
with meat.
002103.9200 We also have a meat pounder for pounding chicken,
or steak, or whatever you want to get to be
002112.4800 more tender.
002113.4800 You might see this also called a tenderizer.
002117.06900 It has some knobs on the top, a tenderizer,
but it's for pounding meat.
002124.0500 All right.
002126.1500 Let's go to the next drawer.
002127.5800 In this drawer we have other random thing.
002131.8700 This is a box of drink accessories.
002135.900 So if you want to make yourself a shot of
some alcohol, if you want to open some wine,
002141.800 we have a wine opener.
002144.3800 There's also a corkscrew.
002147.900 That's this little device is a corkscrew.
002151.4500 So you need to put that into the cork, turn
it, and then you can use the wine opener for
002157.81900 this, or you could ... This is a really old
one.
002161.4700 But you could use a bottle opener to open
beer or cider, whatever you're drinking out
002167.6400 of the bottle.
002168.6400 This is a bottle opener and this is a wine
opener.
002171.500 I also have some wine stoppers because we
don't drink that much wine.
002177.3400 And if we have a bottle of wine, we don't
finish it in one day.
002181.5700 So it's nice to kind of pump the air out.
002187.2700 Our two-year-old likes to play with this a
lot, likes to pump the air out.
002191.7600 And you put the stopper in and it keeps your
wine fresh for longer.
002194.9100 Pretty handy little thing.
002195.9300 We also have a knife sharpener.
002198.5200 This is for our nice, sharp knives, not the
steak knives.
002203.5200 The steak knives, they are serrated.
002207.4800 Serrated means they have little like U marks,
little cuts in them.
002214.2600 Serrated knives are great for cutting tomatoes,
or steak, things like this.
002219.6900 But you cannot sharpen, at least as far as
I know.
002223.60900 You can't sharpen serrated knives, so this
is for our chopping knife or some other knives
002230.5600 that have a smooth blade.
002232.7500 I got this fun little device recently.
002234.57900 Do you know what this is?
002236.7900 Can you guess?
002243.4500 You put a hard boiled egg inside this.
002247.3700 And when you push it down, these little strings
will cut the egg.
002251.08900 I bought this because I saw a toddler using
this to cut hard boiled eggs, and I thought,
002258.6700 "My two-year-old is going to love this."
002261.4600 So I made a bunch of hard boiled eggs and
I peeled them, and he helped me cut them with
002265.7200 this.
002266.8200 It was a hit.
002267.8200 It worked well.
002268.8200 All right.
002269.8200 In this final drawer we have not many options
that you can actually see, but these are Tupperware.
002275.60900 We have some glass Tupperware.
002277.8200 Again, you could say Pyrex dishes.
002281.3200 Technically you could bake, not with the lid,
but you could bake something in this because
002286.3800 it's glass and it's meant for baking.
002289.000 Usually we have a full drawer of these Tupperware
dishes but we don't have a full drawer of
002295.9800 them right now because we've recently been
making a lot of soup.
002300.57900 And as you saw in our freezer, we have a lot
of leftover soup containers.
002305.7700 So when we make soup, maybe chili, we make
a lot, and then put some in the freezer, and
002310.7400 we use those Tupperware containers to store
it.
002314.5400 And then the next week when we want soup,
we don't have to make it.
002317.7700 You just pull it out of the freezer and bam!
002319.8800 You got dinner.
002320.9300 All right.
002321.9300 Let's move on to the next drawer.
002323.4600 These cabinets are under our sink.
002326.7600 So you can see kind of the sink basins are
here.
002329.9900 And usually under the sink is where you keep
extra soap, maybe some cleaning supplies.
002335.6500 We have trash bags, extra trash bags, extra
little cleaning rags.
002341.5400 It's kind of like a cleaning section.
002344.11900 I'm curious what are some things that you
put underneath your kitchen sink.
002350.7100 This is usually where in the US you'll find
cleaning supplies.
002355.32900 We also have a couple appliances.
002356.8400 This is the blender.
002359.6400 We can blend soups.
002361.1800 We can blend smoothies.
002362.1800 You can blend apple sauce, whatever you're
making in there.
002365.5500 And also rice cooker.
002367.04900 Rice cookers are not common kitchen appliances
in the US, but I don't know why not because
002373.4300 they're so handy.
002374.9100 It's so convenient.
002376.1500 We use ours for cooking rice.
002378.400 We use it for steaming some potatoes, things
that take a long time.
002383.1800 And it also doubles as a slow cooker.
002388.4300 So this is a special rice cooker that has
a slow cooker button.
002395.0400 You can just put whatever you want to slow
cook.
002399.3800 I'm not sure if this is a type cuisine that
common in other places but in the US it's
002404.900 really common to cook things in a slow cooker.
002408.0100 A lot of Americans will not have a rice cooker
but they will have a slow cooker.
002412.78900 It looks pretty similar.
002414.57900 Usually it has a glass lid that you can take
off, but it's kind of heavy duty pot.
002420.9800 It's electric, and you can put usually soup
ingredient.
002425.86900 You can put vegetables and broth and you just
push the button.
002430.3200 And after eight hours, usually you go to work,
and you come home, and it's all finished,
002436.4500 and you didn't have to do anything.
002437.9800 It's so great.
002439.0600 So you can use it as a slow cooker.
002440.9600 If you've never heard of a slow cooker, I
recommend looking up some recipes with it.
002447.1700 It's really interesting because slow cookers
have been around for a long time but the idea
002453.86900 of having a good meal cooked without much
time in the kitchen has always been popular.
002462.1900 So a lot of people use slow cookers over all
generations.
002466.5200 It's a pretty cool device to have.
002468.1500 And the next cabinet, pretty nifty.
002471.4200 See how it works.
002472.9500 This opens and, whoa!
002475.2800 This spins.
002476.2800 When we first saw this house, I thought, "Oh,
this so cool.
002480.3100 I can't wait to use this."
002482.34900 The bottom spins, too.
002485.28900 This is just where we keep our dry goods.
002487.61900 So you saw some of the dry goods up on the
shelf.
002490.4100 Those are just kind of some pretty ones for
display.
002492.9800 We do eat those, but it just kind of looks
cool.
002496.0200 In here is where we keep beans, and vinegar,
and tomato paste, and rice, and all of those
002501.9100 things that we use for cooking but we don't
need to keep them on the counter.
002507.4600 A lot of people in the US will have a pantry.
002511.900 This is like a closet in your kitchen where
you can keep all of these types of things.
002518.55900 In this house we don't have a pantry so we
have this lazy Susan is the term.
002525.3900 It's a rolling cabinet.
002528.2400 It could just be a circle that you put a pot
on and that circle spins.
002532.8300 That's called a lazy Susan.
002534.8400 You might see this at restaurants.
002535.9600 You might have it on your kitchen table.
002538.07900 But this is also called a lazy Susan.
002540.3800 I don't know why.
002541.7800 Poor Susan.
002542.7800 I'm sorry if your name's Susan.
002543.82100 You're not lazy.
002545.2500 It's just a term for this type of cool spinning
shelf, and it's really helpful for being able
002553.4900 to access a lot of stuff and make the most
of your space.
002556.31900 We're almost finished.
002557.7500 Hurrah!
002558.7500 These are the last drawers in my kitchen.
002561.400 In these drawers we have our most used kitchen
utensils.
002565.5300 I'd like to go through these because they
might be similar to what you have in your
002568.54900 house.
002570.000 First we have a simple spoon.
002572.10900 This is for stirring things that we're cooking.
002575.61900 This is a garlic press.
002577.5500 Put your garlic in here.
002580.000 When you press it down, the garlic comes out
those holes and it's nice and fine and juicy.
002585.1800 It's really great.
002586.6700 I hate cleaning this, but it's a really handy
device.
002589.4100 Then we have a whisk.
002591.4600 This is a plastic whisk.
002592.9500 If you can see it, it's a plastic whisk, but
a lot of whisks are metal.
002598.58900 They might have a metal part here to it.
002602.7500 This is a soup ladle.
002606.2100 You can ladle.
002607.2200 We can use it as a verb, too.
002608.700 I'm ladling the soup with the soup ladle.
002612.55900 This is a soup ladle.
002614.4100 You might call this too a soup spoon, but
you most likely hear it called a ladle.
002620.9300 Next, this is just a little scraper for our
rice cooker.
002626.1600 This came with it, and I like to scoop up
the rice with it.
002629.1500 This isn't so common in other American households.
002631.35900 But if you have a rice cooker, you probably
have something like this.
002634.6500 We also have a can opener.
002636.4200 This is a manual can opener, so you have to
do this by hand.
002642.0400 When I was growing up, my parents had an electric
can opener, and it was a little device that
002646.56900 stuck on one of our cupboards.
002649.34900 They attached a can to it with a magnet and
just pushed a button and it automatically
002654.2700 opened the can.
002655.35900 Pretty cool.
002656.400 Actually, our cat, the cat that we don't have
anymore, he's 18 years old.
002661.9500 He passed away a couple years ago.
002663.3800 But the cat that I grew up with, he loved
that electric can opener because my dad would
002668.81900 often open cans of tuna with the can opener.
002672.9300 So every time my cat heard the can opener
he would run downstairs and wait.
002678.4400 Please, please, I want tuna.
002680.500 Please.
002681.500 And usually my dad would give him some of
the juice or a little piece or something.
002684.9900 But that sound that electric can opener was
linked in the cat's mind with yummy food.
002690.4500 All right.
002691.4500 Next we have some more tongs.
002692.9800 These are shorter tongs, and a vegetable peeler.
002696.700 I like this kind of vegetable peeler that
has this U shape and the peeler is across
002703.300 the top.
002704.300 But in a lot of houses in the US, you're going
to see a flat vegetable peeler and then the
002710.0100 blade kind of comes out like a knife.
002713.000 I like this kind better.
002714.700 I think it's easier to use, but maybe you
have something different in your house.
002718.8500 We have a spatula.
002720.60900 Actually, we have two spatulas.
002723.81900 This is a silicone spatula and this is a metal
spatula.
002729.9900 We have some more kitchen scissors.
002736.61900 This is another spatula, but we could use
it for stirring or scrapping.
002743.9300 So you might call it a scraper.
002746.4100 You might call it a spatula.
002748.3400 I think most people would probably call it
a spatula even though technically this, this
002754.6300 device that's flat, it's flatter, you can
flip eggs, you can flip things with this spatula,
002762.2300 and this one is more for mixing, like if you're
mixing some bread dough.
002767.59900 Or if you're mixing something, you might be
more likely to use this, but we'd still call
002771.3500 this a spatula.
002772.5200 If I think of another name for this I'll right
it here, but I'm pretty sure we just call
002777.10900 this a spatula.
002778.2100 I also have a lighter.
002780.6500 Hurrah!
002781.6500 Great for lighting candles and whatnot.
002784.4100 Let's go to the next drawer.
002785.6900 Here are all of our spices.
002788.9700 If you watched our other video from our old
apartment, 150 Household Items, that I mentioned
002795.80900 a couple times, in our old apartment, we used
to have all of our spices on a shelf.
002801.9400 Personally, I prefer that because you can
see the names of all the spices and you don't
002807.1300 have to ... All of these, all I see is the
black lid.
002811.1100 So I have to often pick up each one and see
what is this, what is this, what is this.
002816.9100 So I feel like in the future we'll probably
have a more useful way to store our spices.
002824.1800 Maybe we'll put them on a shelf.
002825.8800 But for now, they are in this drawer with
some tea and some other tea filters, other
002831.4200 types of drink things.
002832.7800 But this is where we store our spices.
002835.9200 Then in the bottom drawer ... This is kind
of a weird angle to show you.
002840.61900 But in the bottom drawer are all of our other
things that we need in the kitchen.
002846.61900 So there are small plastic bags.
002850.83900 Sometimes we use the brand name for this,
this type of bag, and we'll call this a Ziploc
002857.6500 bag.
002858.83900 This is not a brand name Ziploc bag, but some
people call it that, even though you might
002866.06900 not actually use that brand.
002868.2300 Ziploc is the brand.
002870.06900 But you might just use a generic or a store
brand name type thing.
002874.85900 So we have some plastic bags.
002876.5600 I have some reusable cloth bags, some reusable,
sealable plastic bags.
002883.7200 I like to use these for cheese because it's
just nice to have something that you can use
002890.2400 again and again.
002892.5500 There's also some extra garbage bags in here.
002896.85900 I'm curious, what do you think this is called
in English?
002901.56900 There are plenty of names.
002905.28900 But in the US we just call this plastic wrap.
002909.1600 Some people might say cling wrap.
002911.9900 I'm sure there are other names for this.
002914.0200 But putting plastic wrap on something, I've
tried a lot of different natural methods for
002920.8900 using instead of plastic wrap, but really,
plastic wrap is so handy.
002926.54900 It's so handy to use.
002928.6600 We also have in here aluminum foil.
002935.7100 Aluminum foil, and we use that whole name,
aluminum foil.
002941.31900 Might put this on a dish when you're baking
it.
002945.3100 You might put it over the dish to keep it
warm.
002947.31900 It's a lot of different uses for that.
002949.5100 I also have some wax paper or parchment paper.
002956.0700 Technically this is not wax paper.
002958.0500 You'll see in the store a different item for
wax paper but some people call it wax paper.
002963.6700 It's technically parchment paper.
002965.80900 But this is what you can put down on your
cookie sheet on your baking sheet so that
002971.9100 you don't have to clean it up that much.
002974.52900 So usually, I like to put this paper on my
baking sheet to just lessen the mess afterwards.
002983.4700 And last but not least are the pots and pans,
essential for cooking.
002992.36900 Mine are hung up really high.
002993.9800 Right now I'm standing on a chair.
002995.83900 These are up here but it's really nice to
be able to have them hung up.
003002.2700 In our kitchen, we don't have a specific place
where we want to put these.
003008.86900 As you can see, all of our drawers are full
of other things so it's really handy to have
003014.6400 these little hooks up here.
003016.0900 If somebody tall comes to our house, they
just have to watch out and not hit their head.
003020.6600 So we hang up our pots, little pots.
003024.82900 These are stainless steel pots, stainless
steel pans, and we also have a nonstick pan.
003035.7400 That's the coating that's on here is going
to help it to not stick.
003040.9700 On this side, this type of ... I'll take it
off so you can see.
003045.9200 This is called a wok.
003048.28900 You're maybe familiar with this type of item.
003051.2600 This is an old wok.
003052.2600 It's supposed to be nonstick but it is quite
old so it's not exactly nonstick anymore.
003057.9600 But, all of these are pots and pans.
003062.07900 You might also hear something like this get
called a sauce pan because it's small.
003069.7100 Some people call that a sauce pan.
003071.07900 I just call it a pot, a small pot, or a pan.
003079.500 So there's simple words.
003080.9100 There's more complex words.
003082.3200 I'm sure there's regional words for these
types of things but I hope that it will help
003085.84900 give you kind of a base for your vocabulary.
003088.4400 Thank you so much for joining me for this
kitchen vocabulary lesson.
003093.9600 It also helped me to have an excuse to clean
the kitchen really well.
003098.7300 My kitchen doesn't usually look like this.
003101.4300 So thank you for giving me that opportunity.
003103.4500 I hope that this lesson helped you to expand
your vocabulary.
003106.2300 And now I want to ask you a question.
003108.4100 What are some kitchen items that I didn't
mention?
003112.04900 Or maybe you have some things in your kitchen
that don't exist in American kitchens.
003117.6400 Let me know in the comments, and make sure
that you read other students' comments as
003120.7300 well because you can just keep expanding your
vocabulary.
003123.77900 Thanks so much for joining me, and I'll see
you again next Friday for a new lesson here
003129.0400 on my YouTube channel.
003130.9700 Bye.
003132.300 The next step is to download my free ebook,
Five Steps to Becoming a Confident English
003138.2500 Speaker.
003139.2500 You'll learn what you need to do to speak
confidently and fluently.
003143.2300 Don't forget to subscribe to my YouTube channel
for more free lessons.
003146.7100 Thanks so much.
003148.05900 Bye.
`
var testdata = "testdata"

func makeDoc(text string) (*Document, error) {
	return NewDocument(
		text,
		WithSegmentation(false),
		WithTagging(false),
		WithExtraction(false))
}

func getTokenText(doc *Document) []string {
	observed := []string{}
	tokens := doc.Tokens()
	for i := range tokens {
		observed = append(observed, tokens[i].Text)
	}
	return observed
}

func getWordData(file string) ([]string, [][]string) {
	in := readDataFile(filepath.Join(testdata, "treebank_sents.json"))
	out := readDataFile(filepath.Join(testdata, file))

	input := []string{}
	output := [][]string{}

	checkError(json.Unmarshal(in, &input))
	checkError(json.Unmarshal(out, &output))

	return input, output
}

func getWordBenchData() []string {
	in := readDataFile(filepath.Join(testdata, "treebank_sents.json"))
	input := []string{}
	checkError(json.Unmarshal(in, &input))
	return input
}

func TestTokenizationEmpty(t *testing.T) {
	doc, _ := makeDoc("")
	assert.Len(t, getTokenText(doc), 0)
}

func TestTokenizationSimple(t *testing.T) {
	doc, _ := makeDoc("Vale is a natural language linter that supports plain text, markup (Markdown, reStructuredText, AsciiDoc, and HTML), and source code comments. Vale doesn't attempt to offer a one-size-fits-all collection of rules—instead, it strives to make customization as easy as possible.")
	expected := []string{
		"Vale", "is", "a", "natural", "language", "linter", "that", "supports",
		"plain", "text", ",", "markup", "(", "Markdown", ",", "reStructuredText",
		",", "AsciiDoc", ",", "and", "HTML", ")", ",", "and", "source",
		"code", "comments", ".", "Vale", "does", "n't", "attempt", "to",
		"offer", "a", "one-size-fits-all", "collection", "of", "rules—instead",
		",", "it", "strives", "to", "make", "customization", "as", "easy", "as",
		"possible", "."}
	assert.Equal(t, expected, getTokenText(doc))
}

func TestTokenizationTreebank(t *testing.T) {
	input, output := getWordData("treebank_words.json")
	for i, s := range input {
		doc, _ := makeDoc(s)
		assert.Equal(t, output[i], getTokenText(doc))
	}
}

func TestTokenizationWeb(t *testing.T) {
	web := `Independent of current body composition, IGF-I levels at 5 yr were significantly
            associated with rate of weight gain between 0-2 yr (beta=0.19; P&lt;0.0005);
            and children who showed postnatal catch-up growth (i.e. those who showed gains in
            weight or length between 0-2 yr by >0.67 SD score) had higher IGF-I levels than other
			children (P=0.02; http://univ.edu.es/study.html) [20-22].`
	expected := []string{"Independent", "of", "current", "body", "composition", ",", "IGF-I",
		"levels", "at", "5", "yr", "were", "significantly", "associated", "with", "rate", "of",
		"weight", "gain", "between", "0-2", "yr", "(", "beta=0.19", ";", "P&lt;0.0005", ")", ";",
		"and", "children", "who", "showed", "postnatal", "catch-up", "growth", "(", "i.e.", "those",
		"who", "showed", "gains", "in", "weight", "or", "length", "between", "0-2", "yr", "by",
		">0.67", "SD", "score", ")", "had", "higher", "IGF-I", "levels", "than", "other", "children",
		"(", "P=0.02", ";", "http://univ.edu.es/study.html", ")", "[", "20-22", "]", "."}
	doc, _ := makeDoc(web)
	assert.Equal(t, expected, getTokenText(doc))
}

func TestTokenizationWebParagraph(t *testing.T) {
	web := `Independent of current body composition, IGF-I levels at 5 yr were significantly
            associated with rate of weight gain between 0-2 yr (beta=0.19; P&lt;0.0005);
            and children who showed postnatal catch-up growth (i.e. those who showed gains in
            weight or length between 0-2 yr by >0.67 SD score) had higher IGF-I levels than other
			children (P=0.02; http://univ.edu.es/study.html) [20-22].

			Independent of current body composition, IGF-I levels at 5 yr were significantly
            associated with rate of weight gain between 0-2 yr (beta=0.19; P&lt;0.0005);
            and children who showed postnatal catch-up growth (i.e. those who showed gains in
            weight or length between 0-2 yr by >0.67 SD score) had higher IGF-I levels than other
			children (P=0.02; http://univ.edu.es/study.html) [20-22].

			Independent of current body composition, IGF-I levels at 5 yr were significantly
            associated with rate of weight gain between 0-2 yr (beta=0.19; P&lt;0.0005);
            and children who showed postnatal catch-up growth (i.e. those who showed gains in
            weight or length between 0-2 yr by >0.67 SD score) had higher IGF-I levels than other
			children (P=0.02; http://univ.edu.es/study.html) [20-22].`

	expected := []string{"Independent", "of", "current", "body", "composition", ",", "IGF-I",
		"levels", "at", "5", "yr", "were", "significantly", "associated", "with", "rate", "of",
		"weight", "gain", "between", "0-2", "yr", "(", "beta=0.19", ";", "P&lt;0.0005", ")", ";",
		"and", "children", "who", "showed", "postnatal", "catch-up", "growth", "(", "i.e.", "those",
		"who", "showed", "gains", "in", "weight", "or", "length", "between", "0-2", "yr", "by",
		">0.67", "SD", "score", ")", "had", "higher", "IGF-I", "levels", "than", "other", "children",
		"(", "P=0.02", ";", "http://univ.edu.es/study.html", ")", "[", "20-22", "]", ".", "Independent", "of", "current", "body", "composition", ",", "IGF-I",
		"levels", "at", "5", "yr", "were", "significantly", "associated", "with", "rate", "of",
		"weight", "gain", "between", "0-2", "yr", "(", "beta=0.19", ";", "P&lt;0.0005", ")", ";",
		"and", "children", "who", "showed", "postnatal", "catch-up", "growth", "(", "i.e.", "those",
		"who", "showed", "gains", "in", "weight", "or", "length", "between", "0-2", "yr", "by",
		">0.67", "SD", "score", ")", "had", "higher", "IGF-I", "levels", "than", "other", "children",
		"(", "P=0.02", ";", "http://univ.edu.es/study.html", ")", "[", "20-22", "]", ".", "Independent", "of", "current", "body", "composition", ",", "IGF-I",
		"levels", "at", "5", "yr", "were", "significantly", "associated", "with", "rate", "of",
		"weight", "gain", "between", "0-2", "yr", "(", "beta=0.19", ";", "P&lt;0.0005", ")", ";",
		"and", "children", "who", "showed", "postnatal", "catch-up", "growth", "(", "i.e.", "those",
		"who", "showed", "gains", "in", "weight", "or", "length", "between", "0-2", "yr", "by",
		">0.67", "SD", "score", ")", "had", "higher", "IGF-I", "levels", "than", "other", "children",
		"(", "P=0.02", ";", "http://univ.edu.es/study.html", ")", "[", "20-22", "]", "."}

	doc, _ := makeDoc(web)
	assert.Equal(t, expected, getTokenText(doc))
}

func TestTokenizationTwitter(t *testing.T) {
	doc, _ := makeDoc("@twitter, what time does it start :-)")
	expected := []string{"@twitter", ",", "what", "time", "does", "it", "start", ":-)"}
	assert.Equal(t, expected, getTokenText(doc))

	doc, _ = makeDoc("Mr. James plays basketball in the N.B.A., do you?")
	expected = []string{
		"Mr.", "James", "plays", "basketball", "in", "the", "N.B.A.", ",",
		"do", "you", "?"}
	assert.Equal(t, expected, getTokenText(doc))

	doc, _ = makeDoc("ˌˌ kill the last letter")
	expected = []string{"ˌˌ", "kill", "the", "last", "letter"}
	assert.Equal(t, expected, getTokenText(doc))

	doc, _ = makeDoc("ˌˌˌ kill the last letter")
	expected = []string{"ˌˌˌ", "kill", "the", "last", "letter"}
	assert.Equal(t, expected, getTokenText(doc))

	doc, _ = makeDoc("March. July. March. June. January.")
	expected = []string{
		"March", ".", "July", ".", "March", ".", "June", ".", "January", "."}
	assert.Equal(t, expected, getTokenText(doc))
}

func BenchmarkTokenization(b *testing.B) {
	in := readDataFile(filepath.Join(testdata, "sherlock.txt"))
	text := string(in)
	for n := 0; n < b.N; n++ {
		_, err := makeDoc(text)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTokenizationSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, s := range getWordBenchData() {
			_, err := makeDoc(s)
			if err != nil {
				panic(err)
			}
		}
	}
}
func TestBigzhu(t *testing.T) {
	//text := `self-assured`
	// text := `fuck—`
	//text := `big:zhu bigzhu/hah bigzhu=zhu bigzhu.com fuck----hah ----big self--assured march...march...... fly…………cry 'fuck so what' I'm your're don't  buy he's wife .`
	text := `& big .`
	//text = `understand...it's`
	// text := `-----fuck`
	// text := `I'm`
	//text := "001.2200 I'm Vanessa from SpeakEnglishWithVanessa.com."

	doc, err := NewDocument(text)
	if err != nil {
		err = errors.WithStack(err)
		t.Fail()
	}

	for _, i := range doc.Tokens() {
		log.Printf("word='%v'", i.Text)
	}
}

func TestReg(t *testing.T) {
	apostropheReg := regexp.MustCompile(`^'.+'`)
	t.Log(apostropheReg.MatchString("'big big'"))
}

func TestSpeed(t *testing.T) {
	start := time.Now()
	doc, err := NewDocument(longArticle)

	if err != nil {
		err = errors.WithStack(err)
		t.Fail()
	}
	doc.Tokens()
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
