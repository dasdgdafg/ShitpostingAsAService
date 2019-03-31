### Triggers
#### !shitpost
Uses a markov chain to generate a high quality post

Put logs from #channel (or whatever else you want to generate shitposts from) in input/#channel, and words you want to not say (such as people you don't want to ping) in avoid/#channel.  You could also just remove the words from the input file and leave the avoid file blank, but I found having two separate files made it easier to change, since I usually end up modifing just one at a time.  Logs from #otherChannel go in input/#otherChannel, etc.  The markov chains are separate per channel because things said in one channel may not be appropriate to post in another.

#### s/foo/bar/
Sed replacements on the most recent matching line

Currently only searches 10 lines back, and returns the first line that matches.  This uses the `--sandbox` flag which was added to sed in 2016, so you may need to update your sed for this to work.

#### !christmas
Tells you how many days until Christmas

Also posts a random link from a text file.  Make a file with the links (to, for example, an image of someone wearing a santa hat), and name it christmas.txt.  

#### !1d6
Roll dice

The number of dice must be in [0, 2^63-1], and the sides per die must be in [0, 2^63-1].  If more than 1000000 dice are rolled, the result is approximated by using a normal distribution.  The dice code uses math/big (\*Float).Sqrt, which is was added in go 1.10.

### Examples
(Using the sample logs for #test)
>\<dasdgdafg> !shitpost  
>\<ShitpostingAsAService> f-oo f-oo bar baz f-oo bar baz bum f-oo f-oo bar bar  
>\<dasdgdafg> !shitpost  
>\<ShitpostingAsAService> bum f-oo f-oo bar bar baz bum  
>\<dasdgdafg> !shitpost  
>\<ShitpostingAsAService> f-oo f-oo f-oo bar baz bum  
>\<dasdgdafg> !shitpost  
>\<ShitpostingAsAService> bum f-oo bar bum f-oo bar baz bum  
  
  
>\<dasdgdafg> foo bar  
>\<dasdgdafg> s/bar/baz/  
>\<ShitpostingAsAService> \<dasdgdafg> foo baz  
>\<dasdgdafg> foo bar  
>\<dasdgdafg> s/bar/baz/  
>\<ShitpostingAsAService> \<dasdgdafg> HAHAHA DISREGARD THAT, I SUCK COCKS  
  
  
>\<dasdgdafg> !christmas  
>\<ShitpostingAsAService> 67 days until Christmas https://files.catbox.moe/9nz3td.png  
  
  
>\<dasdgdafg> !1d20  
>\<ShitpostingAsAService> 3  
>\<dasdgdafg> !3d6  
>\<ShitpostingAsAService> 12  
>\<dasdgdafg> !0d12  
>\<ShitpostingAsAService> 0  
>\<dasdgdafg> !9223372036854775807d9223372036854775807  
>\<ShitpostingAsAService> 42535295872518795703051975483172323328  
