---
title: FAQ
---

# Frequently Asked Questions

[[toc]]

### Why not Kubernetes?

I don't need Kubernetes in my closest. I made Miasma as a little project to host test and IoT applications for my home. Nothing extremely complex, so solutions like Kubernetes are just over the top.

In the end, **Miasma is just crazy simple** and I learned a lot about Docker while working on it. So it was a win-win.

### Why not just use Docker Swarm directly?

I don't want to SSH into my device every time I want to update the app version or change an environment variable. Ideally, apps should just update auto-magically (watchtower plugin incoming).

I also was really bad at managing traefik routing labels by hand, I always seemed to misspell something or mess something else up, meaning my routing was very flaky. Now, Miasma manages routes consistently and more easily then when I did it by hand.

Plus there was inconsistency in more than just routing, `.env` file formatting, `docker-compose.yml` format errors, etc. **Lots of things were inconsistent and now aren't thanks to Miasma**.

### What about Portainer?

Portainer gives you access to every thing Docker can do on a UI, but I wanted to simplify things, most apps don't need all that configuration.

Having miasma manage ports and routing for me is huge, it's so easy to mistype a traefik label and be stuck for hours.
