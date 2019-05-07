package controller

var DEMO bool

const DEMO_TRANSACTIONS = `
[
   {
      "Status":"Confirmed",
      "Date":"08 Jun 17 19:45 +0000",
      "Amount":"37.80251 C",
      "Type":"Transaction",
      "Total":"37.80251 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"08 Jun 17 22:34 +0000",
      "Amount":"-0.11250 C",
      "Type":"Transaction",
      "Total":"37.69001 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"09 Jun 17 18:39 +0000",
      "Amount":"-0.11250 C",
      "Type":"Transaction",
      "Total":"37.57751 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"10 Jun 17 17:28 +0000",
      "Amount":"-0.11250 C",
      "Type":"Transaction",
      "Total":"37.46501 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 20:59 +0000",
      "Amount":"-36.95404 C",
      "Type":"Transaction",
      "Total":"0.51096 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 22:20 +0000",
      "Amount":"40.00000 C",
      "Type":"Transaction",
      "Total":"40.51096 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 22:46 +0000",
      "Amount":"-11.78796 C",
      "Type":"Transaction",
      "Total":"28.72300 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 22:46 +0000",
      "Amount":"-20.85861 C",
      "Type":"Transaction",
      "Total":"7.86439 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 22:46 +0000",
      "Amount":"-4.79004 C",
      "Type":"Transaction",
      "Total":"3.07435 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"290.00000 C",
      "Type":"Transaction",
      "Total":"293.07435 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"-9.34383 C",
      "Type":"Transaction",
      "Total":"283.73052 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"-16.94657 C",
      "Type":"Transaction",
      "Total":"266.78394 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"-3.95982 C",
      "Type":"Transaction",
      "Total":"262.82412 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"-14.05395 C",
      "Type":"Transaction",
      "Total":"248.77017 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-12.53175 C",
      "Type":"Transaction",
      "Total":"236.23842 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-7.35033 C",
      "Type":"Transaction",
      "Total":"228.88809 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-8.73603 C",
      "Type":"Transaction",
      "Total":"220.15206 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-22.81949 C",
      "Type":"Transaction",
      "Total":"197.33257 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-6.80929 C",
      "Type":"Transaction",
      "Total":"190.52328 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-18.47789 C",
      "Type":"Transaction",
      "Total":"172.04539 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-8.83828 C",
      "Type":"Transaction",
      "Total":"163.20710 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-13.24947 C",
      "Type":"Transaction",
      "Total":"149.95763 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-14.45968 C",
      "Type":"Transaction",
      "Total":"135.49795 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-16.32189 C",
      "Type":"Transaction",
      "Total":"119.17606 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-8.28080 C",
      "Type":"Transaction",
      "Total":"110.89526 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-11.23600 C",
      "Type":"Transaction",
      "Total":"99.65926 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-13.57803 C",
      "Type":"Transaction",
      "Total":"86.08123 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-11.23600 C",
      "Type":"Transaction",
      "Total":"74.84524 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-10.36200 C",
      "Type":"Transaction",
      "Total":"64.48324 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-18.18699 C",
      "Type":"Transaction",
      "Total":"46.29625 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-16.74349 C",
      "Type":"Transaction",
      "Total":"29.55276 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-11.09116 C",
      "Type":"Transaction",
      "Total":"18.46160 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-8.94054 C",
      "Type":"Transaction",
      "Total":"9.52106 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-5.44248 C",
      "Type":"Transaction",
      "Total":"4.07858 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"23 Jul 17 21:16 +0000",
      "Amount":"-0.75000 C",
      "Type":"Transaction",
      "Total":"3.32858 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"23 Jul 17 21:21 +0000",
      "Amount":"-0.75000 C",
      "Type":"Transaction",
      "Total":"2.57858 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 05:12 +0000",
      "Amount":"290.00000 C",
      "Type":"Transaction",
      "Total":"292.57858 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-5.36465 C",
      "Type":"Transaction",
      "Total":"287.21394 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-8.60972 C",
      "Type":"Transaction",
      "Total":"278.60422 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-16.90149 C",
      "Type":"Transaction",
      "Total":"261.70273 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-16.90149 C",
      "Type":"Transaction",
      "Total":"244.80124 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-24.85814 C",
      "Type":"Transaction",
      "Total":"219.94310 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-29.74744 C",
      "Type":"Transaction",
      "Total":"190.19566 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-21.55973 C",
      "Type":"Transaction",
      "Total":"168.63594 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-27.57442 C",
      "Type":"Transaction",
      "Total":"141.06152 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-23.48906 C",
      "Type":"Transaction",
      "Total":"117.57246 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-17.89618 C",
      "Type":"Transaction",
      "Total":"99.67628 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-16.90149 C",
      "Type":"Transaction",
      "Total":"82.77479 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-9.91015 C",
      "Type":"Transaction",
      "Total":"72.86464 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-24.58488 C",
      "Type":"Transaction",
      "Total":"48.27976 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-10.93364 C",
      "Type":"Transaction",
      "Total":"37.34612 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-18.66590 C",
      "Type":"Transaction",
      "Total":"18.68022 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-10.65518 C",
      "Type":"Transaction",
      "Total":"8.02504 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-4.98425 C",
      "Type":"Transaction",
      "Total":"3.04079 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:24 +0000",
      "Amount":"58.87511 C",
      "Type":"Transaction",
      "Total":"61.91590 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:59 +0000",
      "Amount":"-18.08436 C",
      "Type":"Transaction",
      "Total":"43.83154 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:59 +0000",
      "Amount":"-17.37911 C",
      "Type":"Transaction",
      "Total":"26.45243 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:59 +0000",
      "Amount":"-16.83380 C",
      "Type":"Transaction",
      "Total":"9.61863 C",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:59 +0000",
      "Amount":"-5.38983 C",
      "Type":"Transaction",
      "Total":"4.22880 C",
      "ID":""
   }
]`
