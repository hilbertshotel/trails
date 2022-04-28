// GET
const get = (id) => {
  return document.getElementById(id)
}

// MAKE
const XXX_MAKEOBJECT_XXX = {
  class: (e, v) => { e.className = v },
  id: (e, v) => { e.id = v },
  text: (e, v) => { e.innerHTML = v },
  value: (e, v) => { e.value = v }
}

const make = (type, ...properties) => {
  const element = document.createElement(type)
  for (const property of properties) {
    const [p1, p2] = property.split("=")
    XXX_MAKEOBJECT_XXX[p1](element, p2)
  }
  return element
}

// INSERT
const insert = (parent, ...elements) => {
  for (const element of elements) {
    parent.append(element)
  }
}

// CLEAR
const clear = (...elements) => {
  for (const element of elements) {
    element.innerHTML = ""
  }
}