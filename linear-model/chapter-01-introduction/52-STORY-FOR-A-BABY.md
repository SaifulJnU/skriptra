# Ch 1 — THE STORY (told to a five-year-old)

> **How to use this file:** read the story once. Then close the file and tell it out loud, in your own words, to nobody. If you can tell it without stopping, you understand Chapter 1. If you get stuck, the place you got stuck is the thing you don't actually understand yet.

---

## The Story of the Tall Grandpa and the Guessing Game

---

Once upon a time there was a little girl named **Nadia**, and Nadia wanted to know something.

She wanted to know: **how tall will my baby brother be when he grows up?**

Her brother was very small. He was mostly a potato. You cannot ask a potato how tall it will be.

So Nadia had an idea.

---

### Part One: Nadia notices something

Nadia went around her whole village and she measured people. She measured all the grown-up children, and she measured their mums and dads.

And she noticed something!

**Tall parents mostly had tall children.**
**Short parents mostly had short children.**

Nadia was very pleased. "I have found the rule!" she said. "Tall makes tall!"

But then she found the Ahmed family. The mum was tall. The dad was tall. And their grown-up son was… **short**.

Nadia was upset. "My rule is broken!"

Her grandma said something very wise. Grandma said:

> **"Your rule is not broken, little one. Your rule is just not the whole story."**
>
> **"There is the part you can guess — and there is the part nobody can guess."**

And that is the **most important sentence in the whole book**.

There is the part you can guess: **tall parents, tall child.**
And there is the part nobody can guess: **sometimes, a surprise.**

Nadia wrote it down:

> **what happens = what I can guess + the surprise**

And the grown-ups have a fancy way of writing exactly the same thing:

$$y = f(x) + \varepsilon$$

The wiggly letter $\varepsilon$ at the end is just the grown-up way of writing **"and then, a surprise."**

---

### Part Two: The surprise is not a mistake

Nadia asked, "Grandma, is the surprise my fault? Did I measure wrong?"

Grandma said: "Sometimes. But mostly, no."

"The surprise is all the things you didn't count. Did the boy eat enough vegetables? Was he sick one winter? Did his great-great-grandmother happen to be short? Was your measuring stick a bit bent?"

"You could never count all of those. Nobody could. So we give all of them **one name**, and we call it **the surprise**, and we put it at the end."

> **The surprise is not a mistake. The surprise is everything you didn't put in your rule.**

And here is the funny bit: **the surprise is not really a thing in the world at all.** If Nadia decides to also count vegetables, then vegetables move out of the surprise and into her rule, and the surprise gets a little smaller. The surprise is the size of *what Nadia doesn't know*, and it changes when Nadia learns more.

---

### Part Three: Where do you put the stick?

Nadia drew all her people as dots on a big piece of paper. Tall parents to the right, short parents to the left. Tall children up high, short children down low.

The dots made a fuzzy sausage shape going up to the right.

"Now," said Nadia, "I will lay a stick across my dots, and the stick will be my guess."

But where do you put the stick? A bit higher? A bit tilted?

Grandma gave her rubber bands.

> "Tie one rubber band from **every single dot** to the stick. Now let go."

The stick wobbled, and settled.

Where did it settle? **Where the rubber bands were pulling least altogether.**

Dots far from the stick pulled HARD. Dots close to the stick barely pulled at all.

So the stick settled in the place where it made **all the dots as un-cross as possible, all at once.**

That is called **least squares**. It's just: *put the stick where the total pulling is smallest.*

And it's why one dot standing very very far away is dangerous — that one dot pulls SO hard that it drags the whole stick toward itself, and now the stick is wrong for everybody else. Grown-ups call that dot an **outlier**, and they worry about it a lot.

---

### Part Four: The thing that made grandma laugh

Nadia looked at her stick and noticed something strange.

The very very tallest parents in the village — their children were tall, yes. But **not quite as tall as the parents.**

And the very very shortest parents — their children were short, but **not quite as short.**

Everyone was sliding a little bit toward the middle!

Nadia panicked. "Is everyone going to become the same size?! Will there be no tall people left?!"

Grandma laughed and laughed.

"No, little one. Think. Why is somebody the *very tallest person in the whole village*?"

Nadia thought.

"…Because they have tall-ness. And also…"

"And also?"

"…and also because they got **lucky**."

"Yes! To be the very tallest, you need tall-ness AND good luck. And here is the thing—"

> **"Tall-ness passes to your children. Luck does not."**

So the children keep the tall-ness and lose the luck, and slide a little toward the middle. Nothing is shrinking. Nothing is broken. **It just looks like a rule, and it isn't one.**

The man who first noticed this was called Galton, and he got so excited that he named the whole subject after it. He called it **regression**.

He was wrong about what it meant. And that is why, to this very day, **the entire subject is named after a mistake.** The grown-ups kept the name to remind themselves: *a pattern can be completely real and still not mean what you think it means.*

---

### Part Five: Nadia gets greedy

Nadia's stick was okay. But it wasn't great. Lots of dots were still far away from it.

"I want a better stick!" she said.

Grandma said, "You have three choices, and only three."

**Choice one: count more things.**
Right now Nadia only counted the parents' height. What if she also counted vegetables, and sleep, and whether they had a big brother? Then more of the surprise moves into the rule. Grown-ups call this **more covariates**.

**Choice two: bend the stick.**
Maybe the dots don't make a straight sausage. Maybe they make a banana! Then a straight stick will never fit, no matter where you put it. You need a **bendy** stick. Grown-ups call this $x^2$, or a log, and — this is the confusing part — **they still call the bendy-stick model a "straight-line model."**

Why?! Because they're not talking about the *stick* being straight. They're talking about **how the numbers get added together**. As long as you're just adding pieces up, they call it linear. Even if the picture is a banana.

**Choice three: give up gracefully.**
Some of the surprise will never go away. Ever. Two identical children in identical houses will still grow to different heights. Chasing that last bit of surprise is chasing nothing.

And there's a trap in chasing it! If Nadia bends her stick SO much that it wiggles through **every single dot perfectly**, she hasn't found the rule — she has just **copied her own piece of paper**. Show her a new child from the next village and her wiggly stick will be terrible.

> **She stopped listening to the music and started copying the crackle.**

Grown-ups call this **overfitting**, and half of this book is about not doing it.

---

### Part Six: Not every question is a stick question

Nadia got so good at sticks that she tried to use one for everything.

"Will it rain tomorrow?" she asked, and laid down her stick.

The stick said: **"1.4 rain."**

Nadia frowned. "What is 1.4 rain? It rains, or it doesn't rain. That's 1 or 0. There is no 1.4."

Then she asked about another day and the stick said **"minus 0.2 rain."**

"That's even worse!"

Grandma nodded. "You cannot carry water in a basket, and you cannot answer a yes-or-no question with a stick."

> **What you're guessing decides what tool you need.**
>
> How tall? → a **stick** (a linear model)
> Yes or no? → a **squashing machine** that never goes below 0 or above 1 (that's called **logit**, and it's the next chapter)
> How many apples? → a **counting box**, because you can't have 2.5 apples (that's called **Poisson**)

And here's the bit people get backwards: **it doesn't matter what you counted to make the guess.** You can count colours, and names, and yes-or-no things. Those are just ingredients; you chop them up and pour them in. **It's only the thing you're guessing *about* that decides which tool.**

---

### Part Seven: The two Nadias

The last thing grandma taught her was the hardest, so listen carefully.

There are **two** Nadias in this story.

There is **the real answer** — the true rule about how heights actually work in the whole wide world, for everybody, forever. Nadia has never seen it. **Nobody has ever seen it.** It's real, it's out there, and it's invisible.

And there is **Nadia's guess** — the stick she made from the 200 people in her own village.

They are not the same thing, and you must never, ever mix them up.

If Nadia walked to the next village and measured 200 different people, she'd get a **slightly different stick**. And the village after that, different again. **Her stick wobbles. The real answer doesn't.**

So grown-ups do something clever. Every time they mean **"Nadia's guess"** instead of **"the real answer,"** they draw a little **hat** on top of the letter.

$$\beta \;\text{(no hat)} = \text{the real answer, invisible, never moves}$$
$$\hat\beta \;\text{(hat)} = \text{Nadia's guess, made from her village, wobbles}$$

That's all a hat means. **"This one came from my village, so be careful."**

And now you can understand the two things grown-ups care about most, because they're both about the hat:

> **"Is my guess pointed at the right place?"** — if you made sticks from a thousand villages and averaged them all, would you land exactly on the real answer? If yes, the guess is **unbiased**. It's aiming true.
>
> **"Is my guess wobbly?"** — how different is the stick from village to village? Less wobble is better. That's **variance**.

Aiming true, and not wobbly. That's the whole dream. And when grown-ups say a guess is **BLUE**, they just mean: *of all the sensible sticks that aim true, this one wobbles the least.*

That's it. That's the whole hat.

---

### The end (and the beginning)

Nadia grew up and became a statistician, and everything she ever did for the rest of her life was one of these five things:

1. 🎵 **Split it up.** What can I guess, and what is surprise? *(y = f(x) + ε)*
2. 🪄 **Lay the stick.** Put it where the rubber bands pull least. *(least squares)*
3. 🎩 **Remember the hat.** My guess is not the real answer. *(β̂ ≠ β)*
4. 🧺 **Check the basket.** Am I guessing a how-much, a yes-or-no, or a how-many?
5. 🚫 **Don't copy the crackle.** A stick that fits my village perfectly will fail the next one.

And the surprise never went away.

And that was fine. **Because knowing exactly how surprised you should be — that turned out to be the most useful thing anyone can know.**

**The End.** 🌙

---

## Your test

Close this file. Now say out loud, in your own words:

1. What are the two parts of everything? *(guessable + surprise)*
2. Where does the stick go, and why? *(least pulling — least squares)*
3. What's in the surprise? *(everything you didn't count — and it shrinks when you count more)*
4. Why did the tall parents have shorter children? *(luck doesn't pass down — and the whole subject is named after this misunderstanding)*
5. What does a hat mean? *(this came from my sample; it wobbles; the real thing doesn't)*
6. Why can't you use a stick for a yes-or-no question? *(it gives 1.4 and −0.2, which aren't probabilities)*
7. What happens if your stick fits every dot perfectly? *(you copied the crackle — overfitting)*

Seven answers, no notes, no stumbling → **Chapter 1 is finished.** Go to Chapter 2.
