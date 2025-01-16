<!-- ISI SCRIPT JS -->
<script setup>
console.info("hello");
</script>

<!-- ISI TAMPILAN BASIC HTML -->
<template>
<h1 class="title">Halaman login</h1>

<div class="card">
  <div class="div-card-body">
    <h5 class="card-title">Halo</h5>
      <p class="card-text">Email</p>
  </div>
</div>
</template>

<!-- ISI STYLING CSS -->
<style scoped>
.title {
  color: black;
  text-align:center ; 
  font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
}

.card-title {
  text-align: center;
  font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
}

.card {
  width: 300px;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.3);
  margin: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 70vh;
  align-items: center;
}

.card-body {
  padding: 20px;
}
</style>