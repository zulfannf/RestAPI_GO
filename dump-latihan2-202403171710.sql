PGDMP  8    
                |            latihan2    16.2    16.2     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16410    latihan2    DATABASE     �   CREATE DATABASE latihan2 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE latihan2;
                postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                pg_database_owner    false            �           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   pg_database_owner    false    4            �            1259    16502    items    TABLE     �   CREATE TABLE public.items (
    item_id integer NOT NULL,
    item_code character varying(50) DEFAULT ''::character varying NOT NULL,
    description text DEFAULT ''::text NOT NULL,
    quantity numeric DEFAULT 0 NOT NULL,
    order_id integer
);
    DROP TABLE public.items;
       public         heap    postgres    false    4            �            1259    16501    items_item_id_seq    SEQUENCE     �   CREATE SEQUENCE public.items_item_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.items_item_id_seq;
       public          postgres    false    218    4            �           0    0    items_item_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.items_item_id_seq OWNED BY public.items.item_id;
          public          postgres    false    217            �            1259    16493    orders    TABLE     �   CREATE TABLE public.orders (
    order_id integer NOT NULL,
    customer_name character varying(255) DEFAULT ''::character varying NOT NULL,
    ordered_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.orders;
       public         heap    postgres    false    4            �            1259    16492    orders_order_id_seq    SEQUENCE     �   CREATE SEQUENCE public.orders_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.orders_order_id_seq;
       public          postgres    false    4    216            �           0    0    orders_order_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.orders_order_id_seq OWNED BY public.orders.order_id;
          public          postgres    false    215            "           2604    16505    items item_id    DEFAULT     n   ALTER TABLE ONLY public.items ALTER COLUMN item_id SET DEFAULT nextval('public.items_item_id_seq'::regclass);
 <   ALTER TABLE public.items ALTER COLUMN item_id DROP DEFAULT;
       public          postgres    false    218    217    218                       2604    16496    orders order_id    DEFAULT     r   ALTER TABLE ONLY public.orders ALTER COLUMN order_id SET DEFAULT nextval('public.orders_order_id_seq'::regclass);
 >   ALTER TABLE public.orders ALTER COLUMN order_id DROP DEFAULT;
       public          postgres    false    215    216    216            �          0    16502    items 
   TABLE DATA                 public          postgres    false    218   (       �          0    16493    orders 
   TABLE DATA                 public          postgres    false    216   �       �           0    0    items_item_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.items_item_id_seq', 4, true);
          public          postgres    false    217            �           0    0    orders_order_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.orders_order_id_seq', 6, true);
          public          postgres    false    215            )           2606    16512    items items_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (item_id);
 :   ALTER TABLE ONLY public.items DROP CONSTRAINT items_pkey;
       public            postgres    false    218            '           2606    16500    orders orders_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public            postgres    false    216            *           2606    16513    items fk_order_id    FK CONSTRAINT     x   ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_order_id FOREIGN KEY (order_id) REFERENCES public.orders(order_id);
 ;   ALTER TABLE ONLY public.items DROP CONSTRAINT fk_order_id;
       public          postgres    false    4647    218    216            �   r   x���v
Q���W((M��L��,I�-Vs�	uV�0�QP762TRŉ�ť9
�@�������5�'!��:��Axd��B lF�~�~SdL��|s�!\\ ��4�      �   �   x���v
Q���W((M��L��/JI-*Vs�	uV�0�QP��ϫT�FF&�ƺ��
�&V�V&fz&f&F������\��3��7��P��@�����h�ƘA�Q�J-*�:���a\\ 8�:@     