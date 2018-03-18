<template>
  <div id="app">
    <nav>
      <div class="nav-wrapper teal">
        <a href="#" class="brand-logo">Jobnetes</a>
        <ul id="nav-mobile" class="right hide-on-med-and-down">
          <li><a href="#">Dummy1</a></li>
          <li><a href="#">Dummy2</a></li>
        </ul>
      </div>
    </nav>
    <div id="main-contents">
      <router-view/>
    </div>
    <footer class="page-footer teal">
      <div class="footer-copyright">
        <div class="container">
          Copyright © 2018 Tatsuya Atsumi Rights Reserved.
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>
