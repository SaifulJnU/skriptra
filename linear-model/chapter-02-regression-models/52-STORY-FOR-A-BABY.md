# Ch 2 — THE STORY (told to a five-year-old)

> Read it once. Then close the file and tell it out loud. Wherever you stumble is the thing you don't understand yet.

---

## The Story of Nadia's Lemonade Stand

---

Nadia is older now. She has a lemonade stand. And she wants to know: **how much money will I make today?**

She has her stick from last time. (The stick that goes where the rubber bands pull least. You remember the stick.)

But today the stick is going to get complicated. Three times.

---

### Part One: The stick that only knows numbers

Nadia's first idea was simple. **The hotter it is, the more lemonade I sell.**

So she wrote down her rule:

> **money = a starting amount + (some amount) × (how hot it is)**

And she laid her stick across her dots, and it worked pretty well!

Then her friend Yusuf came by and said, "Nadia, it's not just the weather. It matters **where** you put your stand."

Nadia said, "Okay! I'll put that in my rule."

And then she got stuck.

Because her stand could be in three places: **by the park**, **by the school**, or **by the station**.

And you cannot multiply a number by "the park."

---

### Part Two: The wall with a mark on it

Nadia's grandma came out with a piece of chalk.

"Come here," she said. "Stand against the wall."

And grandma drew a chalk mark at the top of Nadia's head.

"Now. How tall is Yusuf?"

Nadia said, "One hundred and thirty centimetres."

"No," said grandma. "**Compared to my mark.**"

Nadia looked. Yusuf was a bit taller than the mark. "…Eleven centimetres above the mark."

"Good. And your big cousin?"

"Forty above the mark."

"And **you**?"

Nadia opened her mouth. And then closed it.

Because **she was the mark**. She wasn't above it or below it. She just *was* it.

> **"That's the trick, little one. If you have three places to put your stand, you only need TWO numbers. Because one of the places gets to BE the mark."**

So Nadia picked **the park** to be the mark. And she wrote:

> **money = starting amount + (heat effect) × heat + (school is this much better than the park) + (station is this much better than the park)**

Two extra numbers for three places. **Always one less.**

And the place that became the mark got a special name. Grown-ups call it the **reference category**. Nadia just called it *"the one I'm comparing everything to."*

---

### Part Three: Nadia asks a broken question

The next day Nadia got greedy.

"Why should the park be the mark? That's not fair to the park! I want ALL THREE places to have their own number!"

So she wrote down three numbers, one for each place, **and** she kept her starting amount.

And she gave it to the calculating machine.

And the machine made a horrible noise and **fell over**.

Grandma picked it up. "You broke it."

"How?!"

"You asked it a question with **no answer**. Look."

Grandma wrote on the ground:

> starting amount = 10, park = 0, school = 5, station = 8

"Now watch." She wrote it again:

> starting amount = 0, park = 10, school = 15, station = 18

"Is the school better than the park by five, in both of my answers?"

Nadia checked. "…Yes. Five, both times."

"Is the station better than the park by eight, both times?"

"…Yes."

"So which of my two answers is right?"

Nadia stared. "**They're both right.**"

"They're both right. And so are a million others. And when a question has a million right answers—"

"—it has **no** answer."

> **"That's why one place must be the mark. Not to be unfair. To make the question ASKABLE."**

Grown-ups call this the **dummy variable trap**, and when they say *"X'X is singular"* or *"not full rank,"* they mean exactly this: **the machine fell over because the question had a million answers.**

---

### Part Four: The two escalators

Now Nadia noticed something confusing.

She had two stands. One by the park, one by the station.

On hot days, **both** stands sold more. But the station stand sold *much* more extra on hot days than the park stand did.

Nadia's rule couldn't say that. Her rule said the heat effect was **the same everywhere**, and the station was just *always* a bit better. Like two escalators going up **side by side at the same speed** — one starts higher, and the gap **never changes**.

But that wasn't what was happening. The station escalator was going **faster**.

So Nadia added one more thing to her rule: she **multiplied heat by station**.

> **money = start + (heat effect)×heat + (station bonus) + (extra heat effect if station)×heat×station**

And now the two escalators could go at **different speeds**.

Grown-ups call that last piece an **interaction**, which is a very boring word for a lovely idea:

> **"The size of one thing's effect depends on the other thing."**

And here is the part that trips **everybody** up, so listen:

Once Nadia has the multiplying piece, she can **no longer say "the station bonus is 8."**

Because the station bonus **isn't a number anymore. It's a question.** *A bonus — on what kind of day?* On a cold day, small. On a hot day, huge.

> **When a thing shows up in two places in your rule, you can never again point at one number and call it "the effect."**

That's true for the escalators. And it's true for the ball, too — remember? *"How fast is the ball going?"* has no answer unless you say **when**.

---

### Part Five: The question the stick could not answer

One day a boy walked up and asked Nadia a completely different kind of question.

He said: **"Will it rain tomorrow?"**

Not *how much*. Just — **yes or no**.

Nadia thought, "Easy, I'll use my stick," and she laid the stick down and it said:

> **"1.4 rain."**

Nadia frowned.

She tried another day. The stick said:

> **"minus 0.2 rain."**

"Grandma, my stick is broken."

Grandma shook her head. "Your stick isn't broken. Your stick doesn't know the world has **edges**."

She took Nadia to the bathtub and turned on the tap.

"How full is the tub?"

"Half."

"Now leave the tap on for a hundred years. How full?"

Nadia laughed. "**Full.** Just full. It can't get fuller than full."

"And if I open the drain for a hundred years?"

"**Empty.** It can't get emptier than empty."

"But your stick—" grandma pointed, "—your stick is a **ramp**. A ramp goes up forever and down forever. Your stick has never heard of a floor or a ceiling."

---

### Part Six: The squashing machine

So grandma built Nadia a new machine.

You pour a number in the top — **any** number, huge or tiny, positive or negative — and what comes out the bottom is **always between 0 and 1**. Never above. Never below.

It works like the bathtub:

- Pour in a **huge** number → out comes something like 0.99. Nearly full. But never quite 1.
- Pour in a **very negative** number → out comes 0.01. Nearly empty. But never quite 0.
- Pour in **zero** → out comes exactly 0.5. Right in the middle.

And in the middle is where the machine is **most sensitive**. Near the top and the bottom, you can pour in a LOT and barely anything changes — because the tub is nearly full, or nearly empty, and there's nowhere left to go.

That machine is called the **logit model**. And what you pour in the top is Nadia's old stick — the same stick, unchanged, with all its heat and its places and its escalators. **The stick didn't change. She just put a machine on the end of it.**

---

### Part Seven: The thing Nadia got wrong for a whole week

Nadia built her rain machine. And her rule said:

> **cloudiness gets a 0.7**

So Nadia announced to the whole village: **"Every extra cloud makes rain 0.7 more likely!"**

And it rained, and it didn't rain, and her numbers were nonsense, and she couldn't work out why.

Grandma sat her down.

"Nadia. **Where** does the 0.7 live? At the top of the machine, or the bottom?"

Nadia thought. "…The top. It's part of the stick. It goes **in**."

"And what comes out the bottom?"

"…The chance of rain."

"So is 0.7 a *chance-of-rain* number?"

Nadia went very quiet.

> **"No. It's a going-IN number. And the machine squashes everything on the way through."**

"So one extra cloud adds 0.7 **to what I pour in**. But what that does to the **chance** depends on **where the tub already is.**"

"If the tub is half full — big change! Lots of room to move."
"If the tub is nearly full — tiny change. Almost no room left."

"So there **isn't one number** for how much a cloud changes the chance of rain?"

"There isn't. **There's a different one for every day.**"

Nadia was quiet for a while. Then she said, "So what CAN I say?"

And grandma smiled, because this is the good part.

"You can say it in **gambler's talk**."

"What's gambler's talk?"

"A gambler never says *'70% chance.'* A gambler says *'seven to three on.'* That's called the **odds**. And here's the magic—"

> **"On the odds, your 0.7 works EVERY time, EVERYWHERE, exactly the same."**

"One extra cloud **doubles the odds**. Doubles them if it was already likely. Doubles them if it was very unlikely. **Always doubles.** That's a sentence you can actually say."

And Nadia wrote it on her stand, in big letters, and never got it wrong again:

> ### 🎯 **The number is a GOING-IN number.**
> ### **It doesn't add to the chance. It MULTIPLIES the odds.**

---

### The end

Nadia's stand did very well. And by the end of the summer, her rule had grown up, and it could handle almost anything:

1. 🔢 **Numbers?** Multiply them in.
2. 🧱 **Places, colours, kinds of things?** Pick one to be the **mark**, and count the rest **from** it. Always **one less** than you have.
3. 💥 **Try to give everything its own number?** The machine falls over — because you asked a question with a million answers.
4. 🛗 **Does one thing change how big another thing's effect is?** **Multiply them together.** Now the escalators go at different speeds — and "the effect" isn't a number anymore, it's a question.
5. 🏀 **A thing appearing twice in your rule?** Then it doesn't *have* one effect. Ask **"when?"** or **"for whom?"** first.
6. 🛁 **Is the answer a yes-or-no?** Your stick has no edges. **Put a squashing machine on the end.**
7. 🎲 **And then?** The numbers are **going-in** numbers. Say them in **odds**, not in chances.

And the stick — the plain, ordinary stick from Chapter 1 — **was inside every single one of them.**

That was the secret the whole time. **Nadia never got a new stick. She only ever changed what she poured in, and what she put on the end.**

**The End.** 🍋

---

## Your test

Close the file. Say out loud:

1. Why does Nadia need only **two** numbers for **three** places? *(one place is the mark)*
2. What happens if she gives all three a number? *(a million right answers = no answer; the machine falls over — X'X singular)*
3. What is the "mark" called by grown-ups? *(the reference category)*
4. Yusuf is +11 and cousin is +40 from the mark. How much taller is cousin than Yusuf? *(29 — you subtract; you never use one number alone)*
5. What does multiplying two things together do to the escalators? *(different speeds — non-parallel lines — an interaction)*
6. Why can't Nadia say "the station bonus is 8" once she has an interaction? *(it's not a number anymore, it depends on the heat)*
7. Why does the stick fail for "will it rain?" *(a ramp has no floor and no ceiling; it says 1.4 and −0.2)*
8. Where does the 0.7 live — going in, or coming out? *(going in — it's on the log-odds scale)*
9. What CAN you say about the 0.7? *(it multiplies the odds, the same everywhere)*
10. What never changed in the whole story? *(the stick — the linear predictor x'β)*

Ten answers, no notes → **Chapter 2 is finished.**

Now go to Chapter 3. That's where the stick gets built properly.
