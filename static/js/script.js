const post = () => {}

const elevation = () => {} // grain + loss

const hr = () => {} // avg + max

const pace = () => {} // avrg + best

const distance = () => {} // distance


// DURATION
const duration = (box, data) => {
  // clear box
}

// LOCATION (NAME & TERRAIN)
const loc = (box, data) => {
  const title = make("h3", "text=Location", "id=slide")
  const input = make("input", "id=name")
  const br = make("br")
  
  // dropdown
  const dropdown = make("select")
  const opt1 = make("option", "value=road", "text=road")
  const opt2 = make("option", "value=trail", "text=trail")
  const opt3 = make("option", "value=treadmill", "text=treadmill")
  insert(dropdown, opt1, opt2, opt3)
  
  // output
  const out = make("div", "id=out")
  
  // next button
  const next = make("button", "text=+", "id=next")
  next.onclick = () => {
    // check if input is empty
    // put dropdown + input values in data
    // call duration function
  }

  insert(box, title, input, br, dropdown, out, next)
}

// ON PAGE LOAD
const main = () => {
  const box = get("box")
  loc(box, {})
}

main()