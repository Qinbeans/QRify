<script lang="ts">
  import './assets/styles/base.css'
  import { Stepper, Step } from '@brainandbones/skeleton'
  import { active } from './ts/store'
  import { QRify } from '../wailsjs/go/main/App'

  let size: number = 256
  let url: string = ''
  let name: string = 'default.png'
  const re:RegExp = /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)/

  const complete = () => {
    if(url==='' || !re.test(url)){
      alert('Please enter a valid URL')
      active.set(0)
      return
    }else if(name==='' || name.length>256){
      alert('Please enter a valid name')
      active.set(0)
      return
    }else if(size<1 || size>1024){
      alert('Please enter a valid size')
      active.set(0)
      return
    }
    QRify(url, size, name).catch((err) => {
      console.log(err)
    }).then(() => {
    })
    active.set(0)
  }
  const change = (event: Event) => {
    //get id
    console.log(event)
    let id: string = (<HTMLInputElement>event.target).id 
    switch (id) {
      case "size":
        size = parseInt((<HTMLInputElement>event.target).value)
        if(size > 1024){
          size = 1024
        }else if(size < 0){
          size = 0
        }
        break;
      case "url":
        url = (<HTMLInputElement>event.target).value
        //execute regex
        if(!re.test(url)){
          console.log("invalid url")
          url = ''
        }
        break;
      case "name":
        name = (<HTMLInputElement>event.target).value
        //make sure name is not empty and that it ends with .png
        if(name.length < 1){
          name = 'default.png'
        }else if(!name.endsWith('.png')){
          name = name + '.png'
        }
        break;
      default:
        break;
    }
  }
</script>

<main class="h-screen w-screen bg-gray-900 grid place-content-center text-neutral-900">
  <!-- input link -->
  <!-- input size: default 256  -->
  <!-- input nanme: default.png -->
  <div class="bg-stone-700 p-2 rounded-md">
    <Stepper {active} length={3} on:complete={complete} background="bg-green-500" buttonBack="bg-green-500 rounded px-1 hover:bg-yellow-400 cursor-pointer font-bold" buttonNext="bg-green-500 rounded px-1 {(url=="")?"hover:bg-red-400":"hover:bg-yellow-400 cursor-pointer"} font-bold" buttonComplete="bg-green-600 rounded px-1 hover:bg-yellow-400 cursor-pointer font-bold">
      <Step index={0} locked={url==""}>
        <input class="bg-gray-900 text-white rounded px-1 url" type="text" name="url" id="url" value={url} placeholder="eg https://google.com" on:change={change} required>
      </Step>
      <Step index={1}>
        <input class="bg-gray-900 text-white rounded px-1" type="number" name="size" id="size" value={size} on:change={change}>
      </Step>
      <Step index={2}>
        <input class="bg-gray-900 text-white rounded px-1" type="text" name="name" id="name" value={name} on:change={change} required>
      </Step>
    </Stepper> 
  </div> 
</main>
