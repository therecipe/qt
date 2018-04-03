package controller

var DEMO bool

const DEMO_TRANSACTIONS = `
[
   {
      "Status":"Confirmed",
      "Date":"08 Jun 17 19:45 +0000",
      "Amount":"37.80251 SC",
      "Type":"Transaction",
      "Total":"37.80251 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"08 Jun 17 22:34 +0000",
      "Amount":"-0.11250 SC",
      "Type":"Transaction",
      "Total":"37.69001 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"09 Jun 17 18:39 +0000",
      "Amount":"-0.11250 SC",
      "Type":"Transaction",
      "Total":"37.57751 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"10 Jun 17 17:28 +0000",
      "Amount":"-0.11250 SC",
      "Type":"Transaction",
      "Total":"37.46501 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 20:59 +0000",
      "Amount":"-36.95404 SC",
      "Type":"Transaction",
      "Total":"0.51096 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 22:20 +0000",
      "Amount":"40.00000 SC",
      "Type":"Transaction",
      "Total":"40.51096 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 22:46 +0000",
      "Amount":"-11.78796 SC",
      "Type":"Transaction",
      "Total":"28.72300 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 22:46 +0000",
      "Amount":"-20.85861 SC",
      "Type":"Transaction",
      "Total":"7.86439 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 22:46 +0000",
      "Amount":"-4.79004 SC",
      "Type":"Transaction",
      "Total":"3.07435 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"290.00000 SC",
      "Type":"Transaction",
      "Total":"293.07435 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"-9.34383 SC",
      "Type":"Transaction",
      "Total":"283.73052 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"-16.94657 SC",
      "Type":"Transaction",
      "Total":"266.78394 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"-3.95982 SC",
      "Type":"Transaction",
      "Total":"262.82412 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:30 +0000",
      "Amount":"-14.05395 SC",
      "Type":"Transaction",
      "Total":"248.77017 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-12.53175 SC",
      "Type":"Transaction",
      "Total":"236.23842 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-7.35033 SC",
      "Type":"Transaction",
      "Total":"228.88809 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-8.73603 SC",
      "Type":"Transaction",
      "Total":"220.15206 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-22.81949 SC",
      "Type":"Transaction",
      "Total":"197.33257 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-6.80929 SC",
      "Type":"Transaction",
      "Total":"190.52328 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-18.47789 SC",
      "Type":"Transaction",
      "Total":"172.04539 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-8.83828 SC",
      "Type":"Transaction",
      "Total":"163.20710 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-13.24947 SC",
      "Type":"Transaction",
      "Total":"149.95763 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-14.45968 SC",
      "Type":"Transaction",
      "Total":"135.49795 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-16.32189 SC",
      "Type":"Transaction",
      "Total":"119.17606 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-8.28080 SC",
      "Type":"Transaction",
      "Total":"110.89526 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-11.23600 SC",
      "Type":"Transaction",
      "Total":"99.65926 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-13.57803 SC",
      "Type":"Transaction",
      "Total":"86.08123 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-11.23600 SC",
      "Type":"Transaction",
      "Total":"74.84524 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-10.36200 SC",
      "Type":"Transaction",
      "Total":"64.48324 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-18.18699 SC",
      "Type":"Transaction",
      "Total":"46.29625 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-16.74349 SC",
      "Type":"Transaction",
      "Total":"29.55276 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-11.09116 SC",
      "Type":"Transaction",
      "Total":"18.46160 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-8.94054 SC",
      "Type":"Transaction",
      "Total":"9.52106 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"18 Jul 17 23:42 +0000",
      "Amount":"-5.44248 SC",
      "Type":"Transaction",
      "Total":"4.07858 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"23 Jul 17 21:16 +0000",
      "Amount":"-0.75000 SC",
      "Type":"Transaction",
      "Total":"3.32858 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"23 Jul 17 21:21 +0000",
      "Amount":"-0.75000 SC",
      "Type":"Transaction",
      "Total":"2.57858 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 05:12 +0000",
      "Amount":"290.00000 SC",
      "Type":"Transaction",
      "Total":"292.57858 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-5.36465 SC",
      "Type":"Transaction",
      "Total":"287.21394 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-8.60972 SC",
      "Type":"Transaction",
      "Total":"278.60422 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-16.90149 SC",
      "Type":"Transaction",
      "Total":"261.70273 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-16.90149 SC",
      "Type":"Transaction",
      "Total":"244.80124 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-24.85814 SC",
      "Type":"Transaction",
      "Total":"219.94310 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-29.74744 SC",
      "Type":"Transaction",
      "Total":"190.19566 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-21.55973 SC",
      "Type":"Transaction",
      "Total":"168.63594 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-27.57442 SC",
      "Type":"Transaction",
      "Total":"141.06152 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-23.48906 SC",
      "Type":"Transaction",
      "Total":"117.57246 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"27 Jul 17 23:33 +0000",
      "Amount":"-17.89618 SC",
      "Type":"Transaction",
      "Total":"99.67628 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-16.90149 SC",
      "Type":"Transaction",
      "Total":"82.77479 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-9.91015 SC",
      "Type":"Transaction",
      "Total":"72.86464 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-24.58488 SC",
      "Type":"Transaction",
      "Total":"48.27976 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-10.93364 SC",
      "Type":"Transaction",
      "Total":"37.34612 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-18.66590 SC",
      "Type":"Transaction",
      "Total":"18.68022 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-10.65518 SC",
      "Type":"Transaction",
      "Total":"8.02504 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 00:05 +0000",
      "Amount":"-4.98425 SC",
      "Type":"Transaction",
      "Total":"3.04079 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:24 +0000",
      "Amount":"58.87511 SC",
      "Type":"Transaction",
      "Total":"61.91590 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:59 +0000",
      "Amount":"-18.08436 SC",
      "Type":"Transaction",
      "Total":"43.83154 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:59 +0000",
      "Amount":"-17.37911 SC",
      "Type":"Transaction",
      "Total":"26.45243 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:59 +0000",
      "Amount":"-16.83380 SC",
      "Type":"Transaction",
      "Total":"9.61863 SC",
      "ID":""
   },
   {
      "Status":"Confirmed",
      "Date":"28 Jul 17 18:59 +0000",
      "Amount":"-5.38983 SC",
      "Type":"Transaction",
      "Total":"4.22880 SC",
      "ID":""
   }
]`
