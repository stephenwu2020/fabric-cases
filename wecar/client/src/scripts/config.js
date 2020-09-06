const carList = [
  {
    id: 1,
    brand: "雪弗兰",
    model: "s1",
    seller: "爱车城",
    age: 2,
    price: 100,
  },
  {
    id: 2,
    brand: "福特",
    model: "野马",
    seller: "洛马店",
    age: 3,
    price: 200,
  },
  {
    id: 3,
    brand: "奔驰",
    model: "s200",
    seller: "云海店",
    age: 5,
    price: 300,
  },
  {
    id: 4,
    brand: "奔驰",
    model: "s200",
    seller: "爱车城",
    age: 5,
    price: 300,
  },
  {
    id: 5,
    brand: "奔驰",
    model: "s200",
    seller: "爱车城",
    age: 5,
    price: 300,
  },
  {
    id: 6,
    brand: "奔驰",
    model: "s200",
    seller: "爱车城",
    age: 5,
    price: 300,
  },
  {
    id: 7,
    brand: "奔驰",
    model: "s200",
    seller: "爱车城",
    age: 5,
    price: 300,
  },
  {
    id: 8,
    brand: "奔驰",
    model: "s200",
    seller: "爱车城",
    age: 5,
    price: 300,
  },
  {
    id: 9,
    brand: "奔驰",
    model: "s200",
    seller: "爱车城",
    age: 5,
    price: 300,
  },
  {
    id: 10,
    brand: "奔驰",
    model: "s200",
    seller: "爱车城",
    age: 5,
    price: 300,
  },
]

const carOwner = {
  h_1: [
    { id: 1, time: "2010-12-8", from: "", to: "陈先生" },
    { id: 2, time: "2019-12-8", from: "陈先生", to: "罗女士" }
  ],
  h_2: [
    { id: 1, time: "2005-12-8", from: "", to: "马先生" },
    { id: 2, time: "2019-12-8", from: "马先生", to: "姚女士" }
  ]
}

const carRepair = {
  r_1: [
    { id: 1, time: "2010-5-23", factory: "大平汽修", detail: "修理车灯" },
    { id: 2, time: "2010-8-23", factory: "云海汽修", detail: "汽车保养" },
  ]
}

module.exports = {
  carList,
  carOwner,
  carRepair,
}