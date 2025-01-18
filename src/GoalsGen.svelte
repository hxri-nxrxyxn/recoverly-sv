<script>
    import { Link } from "svelte-routing";
</script>

<main>
    <h1>Generate Effective <br /> <span>Goals</span></h1>
    <div class="gen__box">
        <h3>Short Term Goal</h3>
        <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Eaque dicta
            numquam delectus eligendi ipsam. Alias consectetur assumenda, dicta
            modi suscipit dolor amet? Ea earum blanditiis ipsum rem officiis
            eius asperiores.
        </p>
    </div>
    <div class="btn__box">
        <button class="outlined">REGENERATE</button>
        <br />
        <Link to="/tasks">
            <button class="solid">SAVE</button>
        </Link>
    </div>
</main>

<style>
    .progressi {
        width: 30%;
    }
</style>
