<script>
    import { Link } from "svelte-routing";
    import Nav from "./Nav.svelte";
</script>

<main>
    <Nav />
    <img src="#" alt="" />
    <div class="details">
        <h1>Anti-drug campaign</h1>
        <h2>National Service Scheme</h2>
        <p>
            The NSS students of Thrissur are hosting an Anti-Drug Campaign to
            raise awareness about the dangers of drug abuse and promote a
            drug-free lifestyle. This initiative aims to educate the community
            and offer support to those in need. Join us in spreading the message
            of recovery and building a healthier, drug-free future.
        </p>
        <div class="details__line">
            <h5>Thrissur UNIT 663</h5>
            <p><b>10AM</b> 11 - 09 - 2024</p>
        </div>
    </div>
    <div class="btn__box">
        <Link to="/terms">
            <button class="solid">REGISTER</button>
        </Link>
    </div>
</main>

<style>
    .details__line {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-top: 2rem;
    }
</style>
