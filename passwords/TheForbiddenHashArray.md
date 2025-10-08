# The Forbidden Hash Array: A Mathematical Middle Finger

## TL;DR
This admin password would take longer than the heat death of the universe to crack.<br>
Literally.<br>
We did the math.

## The Setup

Password Requirements:
- Minimum 64 characters
- Truly random (not your cat's name repeated 10 times)
- From the full 94 printable ASCII character set

The Entropy:
```
64 characters × log₂(94) ≈ 64 × 6.555 = 419.52 bits of entropy
```

For context, 256 bits is considered "unbreakable by any conceivable future technology."<br>
We went with 420 bits because I think it is more than 256 bits. Honestly I haven't done the math, but I'm pretty sure it's more.<br>

## Reading The Hash

Let's decode what you're looking at:

```
$2b$17$abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOP
```

Breaking it down:
- `$2b$` = bcrypt algorithm (the `2b` means "bcrypt, current standard implementation")
- `17` = cost factor (this is the important bit)
- Everything after = salt + hash

Cost? Why are you charging me for the hash? Do you accept cashapp?

Bcrypt cost is expressed as 2^n iterations. So:
- Cost 10 = 1,024 iterations
- Cost 12 = 4,096 iterations
- Cost 17 = *131,072 iterations*

Each iteration is intentionally slow. This is a *feature*, not a bug.<br>
No, we don't accept cashapp.

## The Math

Keyspace to search:
```
94^64 ≈ 2^419.52 possible passwords
```

Best-case attacker scenario:
- Bleeding-edge GPU rig
- ~500,000 bcrypt hashes per second at cost 17
- Unlimited budget and electricity
- No laws against theft of global computing resources

Time to crack (50% probability):
```
2^418.52 / 500,000 hashes/sec
≈ 2^399.59 seconds
≈ 1.6 × 10^113 years
```

For perspective:
- Age of the universe: ~1.4 × 10^10 years
- Estimated time until heat death: ~10^100 years (we're in the ballpark!)
- Number of atoms in the observable universe: ~10^80

You would need to harness the energy output of *every star in the observable universe* and:
1. Convert it all into bcrypt ASICs at 100% efficiency
2. Run them continuously...
3. ??
3. Profit? (Not quite, you're still a couple orders of magnitude short)

## The Monthly Rotation

Because apparently 10^113 wasn't overkill enough, these hashes rotate *automatically every month* for 10 years.

How it works:
1. Script generates 120 unique bcrypt hashes in advance (one per month, 2025-2034)
2. Auth system checks current year-month and uses corresponding hash
3. Each month = different hash = even if you somehow cracked one (you won't), you get 30 days of access before it rotates

"Well I'll just crack the last one and then wait a decade to use the password!"<br>
Lol fair enough, you win.

Dude, why?
- Defense in depth? Sure, let's go with that.
- Because it's hilarious? Definitely.
- Because we can? Absolutely.

This isn't just security. This is security *performance art*.

I wanted to break security best practices to bad I had to use numbers that have no physical meaning.<br>
The universe will experience heat death. Protons will decay. The last black holes will evaporate via Hawking radiation.<br>
But this password will remain.

---

*"But why not just use environment variables?"*  
What are environment variables? Never heard of them.

*"Isn't this overkill?"*  
Brother, we passed overkill at the 32-character mark. Now we're just style points.

*"What if quantum computers—"*  
Even Shor's algorithm doesn't help with brute-forcing symmetric keys.<br>
You'd need Grover's algorithm, which would "only" square-root the search time.<br>
That brings us down to... *checks notes* ...still 10^42 years.

*"I'm going to get the password out of the passwords.txt file on your desktop."*
Nuh uh, I wrote it on a sticky note an... Oh, you almost got me there you rascal you!