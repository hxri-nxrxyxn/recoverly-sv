<script>
    import Nav from "./Nav.svelte";
    import { Link } from "svelte-routing";
    import ProfilePic from "./assets/circle-user.svg";
</script>

<main>
    <Nav />
    <h3>Profile</h3>
    <div class="profile__box">
        <div class="profile__image">
            <img src={ProfilePic} alt="" class="icon" />
        </div>
        <div class="profile__text">
            <p>
                <b>Hari Narayan M R</b>
                <br />
                20, Male
            </p>
        </div>
    </div>
    <div class="bio__box">
        <h3>Bio</h3>
        <p>
            Lorem ipsum, dolor sit amet consectetur adipisicing elit. Numquam
            voluptatibus distinctio sint dignissimos, reiciendis perspiciatis
            necessitatibus modi. Animi eos reiciendis praesentium nostrum
            tempora cumque eveniet, corrupti dolores nesciunt sequi, asperiores
            dolorum dolor atque voluptatem excepturi. Velit modi ut illum fuga
            repudiandae tenetur sed molestiae ratione est. Iste quaerat dolor
            facere dolorem nihil rem quos aliquid quae excepturi. Soluta eaque
            minima vel ipsum quisquam iste quis nihil deleniti magnam! Omnis
            quas cum illum esse necessitatibus inventore eos et, odit explicabo
            voluptatum quibusdam perferendis dicta tempore hic eligendi neque
            similique! Dicta, quod!
        </p>
    </div>
</main>

<style>
    .scan__cam {
        height: 50vh;
        background-color: var(--color-light);
        margin: 0.5rem 0 2rem 0;
    }
    .profile__box {
        margin: 2rem 0;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    .profile__image {
        width: 40%;
    }
    .profile__image img {
        width: 100%;
    }
</style>
