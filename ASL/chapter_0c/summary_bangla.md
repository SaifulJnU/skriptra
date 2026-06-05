# 📊 অধ্যায় ০c: Exploratory Data Analysis (EDA) — বাংলায় সম্পূর্ণ ব্যাখ্যা

> 📝 **নোট:** *এইটা কোর্সের official অধ্যায় নয় — এইটা আমি (Saiful) আমার নিজের বোঝার জন্য add করেছি।* EDA হলো ML/Statistics-এর ভিত্তি, তাই আলাদা করে গুছিয়ে রাখলাম।

---

# EDA আসলে কী?

🔊 **উচ্চারণ:** *Exploratory Data Analysis* → "এক্সপ্লোরেটরি ডেটা অ্যানালাইসিস"।

**বাংলায়:** "অনুসন্ধানমূলক ডেটা বিশ্লেষণ" বা "প্রাথমিক ডেটা অনুসন্ধান ও বিশ্লেষণ"।

👶 **সহজ ভাষায়:** EDA হলো কোনো ডেটাসেটকে **মডেল বানানোর আগে** ভালোভাবে দেখে বোঝার প্রক্রিয়া। ডেটার মধ্যে কী তথ্য আছে, কোনো ভুল বা অস্বাভাবিক মান (outlier) আছে কি না, ডেটার প্যাটার্ন বা সম্পর্ক কী — এসব খুঁজে বের করা।

> 🎯 **এক লাইনে:** EDA = ডেটা নিয়ে কোনো মডেল বানানোর **আগে** ডেটাকে বুঝতে, পরিষ্কার করতে এবং এর ভেতরের প্যাটার্ন খুঁজে বের করার প্রক্রিয়া।

---

# EDA-তে সাধারণত যা করা হয়

- ডেটার সারাংশ দেখা (Mean, Median, Standard Deviation)
- Missing values আছে কি না দেখা
- Outliers খুঁজে বের করা
- Distribution দেখা (ডেটা কীভাবে ছড়িয়ে আছে)
- Variables-এর মধ্যে সম্পর্ক দেখা
- Graphs ও Charts ব্যবহার করা (Histogram, Box Plot, Scatter Plot)

---

# 🎓 উদাহরণ: বিশ্ববিদ্যালয়ের GPA ডেটা

ধরো একটি বিশ্ববিদ্যালয়ের শিক্ষার্থীদের GPA-এর ডেটা আছে। EDA করলে তুমি:

- **গড় GPA** কত তা দেখবে। (যেমন mean = 3.2)
- সবচেয়ে বেশি GPA **কোন রেঞ্জে** আছে তা দেখবে। (যেমন বেশিরভাগ 3.0–3.5)
- কোনো **অস্বাভাবিক GPA** (যেমন 15.0, যা ভুল — কারণ GPA সর্বোচ্চ 4.0) আছে কি না খুঁজবে → **outlier**।
- **GPA ও Study Hours**-এর মধ্যে সম্পর্ক আছে কি না দেখবে → বেশি পড়লে কি GPA বাড়ে? (**correlation**)

---

# 🗺️ EDA Mind Map (পুরো কাঠামো)

> ক্লাসে কোনো শব্দ শুনলে যেন সঙ্গে সঙ্গে বুঝতে পারো এটা কোন অংশের — তাই পুরো EDA-কে গাছ আকারে সাজানো হলো।

```
Exploratory Data Analysis (EDA)
│
├── 1. Data Understanding (ডেটা বোঝা)
│   ├── Dataset
│   ├── Observation / Sample
│   ├── Variable / Feature
│   ├── Population
│   └── Data Types
│       ├── Numerical
│       │   ├── Discrete
│       │   └── Continuous
│       └── Categorical
│           ├── Nominal
│           └── Ordinal
│
├── 2. Data Cleaning (ডেটা পরিষ্কার করা)
│   ├── Missing Values
│   ├── Duplicate Data
│   ├── Outliers
│   ├── Data Entry Errors
│   └── Inconsistent Values
│
├── 3. Univariate Analysis (একটি ভেরিয়েবল)
│   ├── Purpose: একটি ভেরিয়েবলের বৈশিষ্ট্য বোঝা
│   ├── Statistics
│   │   ├── Mean (গড়)
│   │   ├── Median (মধ্যক)
│   │   ├── Mode (বহুলক)
│   │   ├── Range
│   │   ├── Variance
│   │   └── Standard Deviation
│   └── Visualization
│       ├── Histogram
│       ├── Box Plot
│       ├── Density Plot
│       ├── Bar Chart
│       └── Pie Chart
│
├── 4. Bivariate Analysis (দুটি ভেরিয়েবল)
│   ├── Purpose: দুটি ভেরিয়েবলের সম্পর্ক দেখা
│   ├── Numerical vs Numerical
│   │   ├── Correlation
│   │   └── Scatter Plot
│   ├── Categorical vs Numerical
│   │   ├── Group Comparison
│   │   ├── Box Plot
│   │   └── Violin Plot
│   └── Categorical vs Categorical
│       ├── Contingency Table
│       ├── Cross Tabulation
│       └── Chi-Square Test
│
├── 5. Multivariate Analysis (৩+ ভেরিয়েবল)
│   ├── Purpose: একাধিক ভেরিয়েবলের যৌথ সম্পর্ক দেখা
│   ├── Correlation Matrix
│   ├── Pair Plot
│   ├── Heatmap
│   ├── PCA
│   ├── Clustering
│   └── Feature Interaction
│
├── 6. Probability & Distribution
│   ├── Normal Distribution
│   ├── Skewness
│   ├── Kurtosis
│   ├── Probability
│   ├── Sampling Distribution
│   └── Central Limit Theorem
│
├── 7. Statistical Inference
│   ├── Population
│   ├── Sample
│   ├── Estimation
│   ├── Confidence Interval
│   ├── Hypothesis Testing
│   ├── p-value
│   └── Significance Level
│
├── 8. Predictive Analysis Preparation
│   ├── Feature Engineering
│   ├── Feature Selection
│   ├── Data Transformation
│   ├── Scaling
│   ├── Normalization
│   └── Encoding
│
└── 9. Common Classroom Terms
    ├── Variable / Feature
    ├── Target Variable
    ├── Independent / Dependent Variable
    ├── Correlation / Covariance
    ├── Distribution
    ├── Bias / Variance
    ├── Overfitting / Underfitting
    ├── Sampling
    ├── Random Variable
    └── Model
```

---

# 🌟 সবচেয়ে গুরুত্বপূর্ণ অংশ: ৩টি বিশ্লেষণ গ্রুপ

> EDA মনে রাখার সবচেয়ে সহজ উপায় — **কয়টা variable একসাথে দেখছি** তা দিয়ে ভাগ করা।

## ① Univariate — "একটা ভেরিয়েবল দেখি"
🔊 *ইউনিভেরিয়েট* (uni = এক)।
📖 **উদ্দেশ্য:** একটি variable-এর নিজস্ব বৈশিষ্ট্য বোঝা (center কোথায়, কতটা ছড়ানো, আকৃতি কেমন)।
🛠️ **টুল:** Mean, Median, Mode, Variance, SD · Histogram, Box Plot, Density/Bar/Pie।
✍️ **উদাহরণ:** শুধু GPA-এর histogram — বেশিরভাগ ছাত্র কোন GPA-তে আছে?

## ② Bivariate — "দুইটা ভেরিয়েবলের সম্পর্ক দেখি"
🔊 *বাইভেরিয়েট* (bi = দুই)।
📖 **উদ্দেশ্য:** দুটি variable-এর মধ্যে সম্পর্ক আছে কি না দেখা। তিন রকম জোড়া:
- **Numerical vs Numerical:** Correlation, Scatter Plot। ✍️ Study Hours vs GPA।
- **Categorical vs Numerical:** Group comparison, Box/Violin Plot। ✍️ বিভাগ (CSE/EEE) vs GPA।
- **Categorical vs Categorical:** Contingency Table, Cross-tab, Chi-Square। ✍️ লিঙ্গ vs পাস/ফেল।

## ③ Multivariate — "অনেকগুলো ভেরিয়েবল একসাথে দেখি"
🔊 *মাল্টিভেরিয়েট* (multi = বহু, ৩+)।
📖 **উদ্দেশ্য:** একাধিক variable-এর যৌথ সম্পর্ক ও প্যাটার্ন।
🛠️ **টুল:** Correlation Matrix, Pair Plot, Heatmap, PCA, Clustering, Feature Interaction।
✍️ **উদাহরণ:** GPA + Study Hours + Attendance + ঘুম — সব একসাথে heatmap-এ দেখা।

> 🧠 **মনে রাখার ট্রিক:**
> - **Uni → একটা** (Mean, Median, Histogram, Boxplot)
> - **Bi → দুইটার সম্পর্ক** (Correlation, Scatter Plot)
> - **Multi → অনেকগুলো** (Heatmap, PCA, Clustering)

---

# 📖 ক্লাসে শোনার মতো গুরুত্বপূর্ণ শব্দ (অভিধান)

| English | বাংলা | মানে (এক লাইনে) |
|---------|-------|------------------|
| Dataset | ডেটাসেট / তথ্যসমষ্টি | পুরো তথ্যের সংগ্রহ |
| Variable / Feature | চলক / বৈশিষ্ট্য | একটা কলাম (যেমন GPA) |
| Observation / Sample | নমুনা / পর্যবেক্ষণ | একটা সারি (একজন ছাত্র) |
| Population | জনসংখ্যা / সমগ্রক | পুরো গোষ্ঠী (সব ছাত্র) |
| Distribution | বণ্টন | ডেটা কীভাবে ছড়িয়ে আছে |
| Mean | গড় | সব মান যোগ ÷ সংখ্যা |
| Median | মধ্যক | মাঝখানের মান |
| Mode | বহুলক | সবচেয়ে বেশিবার আসা মান |
| Variance | বিচ্যুতি/ভেদাঙ্ক | ছড়ানোর পরিমাপ (বর্গ) |
| Standard Deviation | মানক বিচ্যুতি | variance-এর বর্গমূল |
| Outlier | অস্বাভাবিক মান | বাকিদের থেকে অনেক দূরে |
| Correlation | পারস্পরিক সম্পর্ক | দুটো একসাথে ওঠে/নামে কি না |
| Covariance | সহভেদাঙ্ক | correlation-এর scale-বিহীন রূপ |
| Histogram | হিস্টোগ্রাম | bar দিয়ে distribution |
| Box Plot | বক্স প্লট | median + ছড়ানো + outlier |
| Scatter Plot | স্ক্যাটার প্লট | দুই সংখ্যার বিন্দু-গ্রাফ |
| Skewness | বঙ্কিমতা | distribution একপাশে কাত |
| Kurtosis | কুর্টোসিস | চূড়া কত তীক্ষ্ণ |
| Hypothesis Testing | প্রকল্প যাচাই | দাবি সত্য কি না পরীক্ষা |
| p-value | পি-মান | দৈবচয়নে এমন ফল আসার সম্ভাবনা |
| Confidence Interval | আস্থা ব্যবধান | সত্য মান যে পরিসরে থাকার সম্ভাবনা |

---

# 🧩 ছোট উদাহরণ-অঙ্ক

### সমস্যা ১ — Univariate (Mean ও Median)
৫ জন ছাত্রের GPA: {2.0, 3.0, 3.0, 3.5, 4.0}। Mean ও Median বের করো।

**সমাধান:** Mean = (2.0+3.0+3.0+3.5+4.0)/5 = 15.5/5 = **3.1**। সাজানো মাঝখানের মান (৩য়) = **3.0** = Median।

### সমস্যা ২ — Outlier চেনা
GPA ডেটায় একটা মান **15.0**। এটা কী এবং কেন?

**সমাধান:** এটা **outlier** (আসলে data-entry error) — কারণ GPA সর্বোচ্চ 4.0, তাই 15.0 অসম্ভব। Data Cleaning ধাপে এটা ঠিক বা বাদ দিতে হবে।

### সমস্যা ৩ — কোন analysis?
(ক) শুধু GPA-এর histogram। (খ) Study Hours vs GPA-এর scatter plot। (গ) GPA + Attendance + ঘুম-এর heatmap। প্রতিটা কোন গ্রুপ?

**সমাধান:** (ক) **Univariate** (১টা variable)। (খ) **Bivariate** (২টা)। (গ) **Multivariate** (৩+)।

---

> 🎯 **শেষ কথা:** ক্লাসে **Distribution, Sampling, Confidence Interval, Correlation, Variance, Standard Deviation, Outlier, Hypothesis Testing** — এই শব্দগুলো শুনলেই বুঝবে এগুলো EDA ও Statistics-এর সবচেয়ে গুরুত্বপূর্ণ ভিত্তি। আর সবসময় মনে রেখো: **Uni → একটা, Bi → দুইটা, Multi → অনেকগুলো।** 📊
